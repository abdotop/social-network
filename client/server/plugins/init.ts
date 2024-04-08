export default defineNitroPlugin(() => {
    runTask('socket:connect')
    runTask('notif:sendNotif')
    runTask('message:private_message')
})
