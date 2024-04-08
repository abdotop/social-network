import type { Peer } from "crossws";
import { getQuery } from "ufo";
import { server_socket } from "~/stores/socketCon"

const notifconns = new Map<string, { conn: Peer, onLine: boolean }>();
const messageconns = new Map<string, { conn: Peer, otherId: string }>();

export const notifUser = (data: any) => {
  const user = notifconns.get(data.concernID)
  console.log(data.type ,data.concernID, user, "notifUser");

  if (user) {
    console.log(user.onLine, "notifUser");
    
    if (user?.onLine) {
      if (data.type === 'new_message') {
        const otherId = messageconns.get(data.concernID)?.otherId
        // console.log(otherId, data.user.ID, "online message");
        if (!otherId || otherId !== data.user.ID) {
          user.conn.send(data)
        }
      } else {
        user.conn.send(data)
      }
    }
  }
}

export const messageUser = (data: any) => {

  const user = messageconns.get(data.ReceiverID)
  if (user) {
    user.conn.send({ type: 'new_message', data })
  }
  const sender = messageconns.get(data.SenderID)
  if (sender) {
    sender.conn.send({ type: 'new_message', data })
  }
}

export default defineWebSocketHandler({
  async open(peer: Peer) {
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    if (channel === "notif") {
      console.log(userId, "notif");
      
      notifconns.set(userId, { conn: peer, onLine: true });
    }
    if (channel === "message") {
      const otherId = query.otherId as string;
      messageconns.set(userId, { conn: peer, otherId });
    }
  },
  async message(peer: Peer, message) {
    const data = JSON.parse(message.toString())
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    // if (channel === "notif") {
    //   notifUser(data)
    // }
    if (channel === "message") {
      if (data.type === 'online') {
        const users = Array.from(data.users).map((user: any) => {
          const online = notifconns.get(user.id)?.onLine
          if (online) {
            return { ...user, online: true }
          }
          return user
        })
        peer.send(JSON.stringify({ type: 'online', users }))
      } else if (data.type === 'private_message') {
        server_socket!.send(data)
      }
    }
  },

  close(peer: Peer, details) {
    console.log(`[ws] close ${peer}`);
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    if (channel === "notif") {
      notifconns.set(userId, { conn: peer, onLine: false });
    } else if (channel === "message") {
      messageconns.delete(userId);
    }



    // const userId = getUserId(peer) /;
    // users.set(userId, { online: false });
  },

  // error(peer: Peer, error) {
  //   console.log(`[ws] error ${peer}`, error);
  // },

  upgrade(req) {
    return {
      headers: {
        "x-powered-by": "cross-ws",
      },
    };
  },
});

function get(peer: Peer) {
  const query = getQuery(peer.url);
  return query
}

// function getStats() {
//   const online = Array.from(users.values()).filter((u) => u.online).length;
//   return { online, total: users.size };
// }
