export const getNotifications = async () => {
    const res = await fetch("/api/notification/getnotification").then()
    const responseInJsonFormat = await res.json().catch(err => ({ error: err }))

    return responseInJsonFormat
}