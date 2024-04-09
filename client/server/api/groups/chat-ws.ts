import type { Peer } from "crossws";
import { getQuery } from "ufo";
import type { GroupMessage, ServerResponse, UserWithoutPassword } from "~/types";

interface GroupMember {
  id: string;
}

const groups = new Map<string, GroupMember[]>();

export default defineWebSocketHandler({
  async open(peer) {
    const config = useRuntimeConfig();
    // @ts-ignore
    const _cookie = decodeURIComponent(peer.headers.cookie.split(config.cookieName + '=').pop().split(';')[0]);
    console.log(`[ws] cookie ${_cookie}`);
    if (!_cookie) {
      return;
    }
    const { user, token } = await getUserFromToken(_cookie);
    if (!user || !token) {
      return;
    }

    const groupId = getGroupId(peer);
    const group = groups.get(groupId);
    if (!group) {
      groups.set(groupId, [{ id: user.id }]);
    } else {
      group.push({ id: user.id });
      groups.set(groupId, group);
    }

    peer.send({
      Sender: { nickname: "server" },
      Content: `Welcome to the server ${groupId}!`,
    });

    peer.subscribe("chat" + groupId);
    // peer.publish("chat" + groupId, { Sender: { nickname: "server" }, Content: `${user.nickname} joined!` });
  },

  async message(peer, message) {
    console.log(`[ws] message ${message.text()}`);

    const groupId = getGroupId(peer);
    if (message.text() === "ping") {
      peer.send({ user: "server", message: "pong" });
      return
    }

    const config = useRuntimeConfig();
    // @ts-ignore
    const _cookie = decodeURIComponent(peer.headers.cookie.split(config.cookieName + '=').pop().split(';')[0]);
    console.log(`[ws] cookie ${_cookie}`);
    if (!_cookie) {
      return;
    }
    const { user, token } = await getUserFromToken(_cookie);
    if (!user || !token) {
      return;
    }

    // Store message in database
    const _response = await $fetch(`${process.env.BACKEND_URL}/group/messages/new?group_id=${groupId}`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        content: message.text(),
      }),
    });

    if (!_response) {
      return;
    }
    const response = JSON.parse(_response as string) as ServerResponse<GroupMessage>;

    const { password: _password, ...userWithoutPassword } = user;
    console.log(`[ws] user ${JSON.stringify(userWithoutPassword)}`);
    const _message = {
      ...response.data,
      Sender: userWithoutPassword,
    } as GroupMessage;
    console.log(`[ws] message ${_message.Content}`);
    peer.send(_message); // echo back
    peer.publish("chat" + groupId, _message);
  },

  close(peer, details) {
    console.log(`[ws] close ${peer}`);

    const groupId = getGroupId(peer);
    const group = groups.get(groupId);
    if (!group) {
      return;
    }
    // remove from group set
    const index = group.findIndex((member) => member.id === peer.id);
    if (index > -1) {
      group.splice(index, 1);
    }
    groups.set(groupId, group);
    // peer.publish("chat" + groupId, { Sender: { nickname: "server" }, Content: `${peer.id} left!` });
  },

  error(peer, error) {
    console.log(`[ws] error ${peer}`, error);
  },

  upgrade(req) {
    return {
      headers: {
        "x-powered-by": "cross-ws",
      },
    };
  },
});

function getGroupId(peer: Peer) {
  const query = getQuery(peer.url);
  return query.group_id as string;
}
