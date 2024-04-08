import type { Group, GroupMember, Invitation, ServerResponse } from '~/types';

export const useGroupRequest = () => {
    async function joinRequest(groupId: string | undefined): Promise<{ data: any, error: any }> {

        const { data, error } = await useFetch("/api/groups/request/join", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
            },
        });

        return { data, error };
    }

    async function getJoinRequests(
        groupId: string | undefined
    ): Promise<GroupMember[] | null> {

        const response = await $fetch("/api/groups/request/join-requests", {
            method: "GET",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
            },
        });
        // console.log('###################################\n',data,'###################################\n');

        const data = JSON.parse(response as string)


        return data.data as GroupMember[]
    }

    async function acceptJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {
        
        const data = await $fetch("/api/groups/request/accept", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gId,
                rId,
            },
        });
        return { data };
    }

    async function declineJoinRequest(
        gId: string,
        rId: string
    ): Promise<any> {

        const data = await $fetch("/api/groups/request/decline", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gId,
                rId,
            },
        });
        return { data };
    }

    async function getUserGroups() {

        const response = await $fetch("/api/user/groups", {
            method: "GET",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

        });

        return response
    }

    async function sendInvitation(groupId: string, userId: string) {
        const { data, error } = await useFetch("/api/groups/request/sendInvitation", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                gid: groupId,
                uid: userId
            },
        });

        return { data, error };
    }

    async function acceptInvitation(memberId: string) {
        const { data, error } = await useFetch("/api/groups/invitations/accept", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                iId: memberId
            },
        });

        return { data, error };
    }

    async function declineInvitation(memberId: string) {
        const { data, error } = await useFetch("/api/groups/invitations/reject", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,

            query: {
                iId: memberId
            },
        });

        return { data, error };
    }


    async function getInvitations() {
        const response = await $fetch<string>("/api/groups/invitations", {
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
        });

        console.log(response);
        

        return JSON.parse(response) as ServerResponse<Invitation[]>;
    }

    return {
        joinRequest,
        getUserGroups,
        declineJoinRequest,
        acceptJoinRequest,
        getJoinRequests,
        sendInvitation,
        acceptInvitation,
        declineInvitation,
        getInvitations
    }

}