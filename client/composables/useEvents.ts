import type { Event, EventParticipant, ServerResponse } from "~/types";

export const useEvents = () => {
    async function createEvent(eventDetail: any, groupId: string) {
        const { data, error } = await useFetch("/api/groups/events/create", {
            method: "post",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            body: JSON.stringify(eventDetail),
            query: {
                gid: groupId,
            }
        })

        if (error.value) {
            return { error }
        }
        console.log('created event bbbb\n', data.value);

        const resp = JSON.parse(data.value as string) as ServerResponse<Event>
        console.log('created event\n', resp, error);

        return { data: resp.data }
    }

    const getAllEvents = async (groupId: string) => {
        return await useAsyncData(
            'events',
            () => $fetch<any>("/api/groups/events", {
                headers: useRequestHeaders(["cookie"]) as HeadersInit,
                query: {
                    participants: 1,
                    user: 1,
                    gid: groupId
                }
            })
        )
    };

    async function respondEvent(event: Event, response: string) {
        const { data, error } = await useFetch("/api/groups/events/respond", {
            method: "POST",
            headers: useRequestHeaders(["cookie"]) as HeadersInit,
            query: {
                eid: event.ID,
                gid: event.GroupID
            },
            body: JSON.stringify({ response })
        });

        console.log(data);


        return { data: JSON.parse(data.value as unknown as string) as { data: any, message: string }, error };
    }

    return { createEvent, getAllEvents, respondEvent }
}