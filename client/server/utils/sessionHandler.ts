import { useSession, H3Event, sendError, getSession } from 'h3'

export const sessionCreator = async (token: any, user: any, event: H3Event) => {
  if (token != undefined) {

  }
  const serverSession = await useSession(event, {
    // PASSWORD TO change as .env variable
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
    cookie: {
      httpOnly: true,
      secure: true,
      sameSite: "strict",
    },
    maxAge: 60 * 60 * 24 * 7,
    generateId: () => { return token }
  });

  await serverSession.update({
    userInfos: user,
    sessionToken: token
  });

  return serverSession
}

export const sessionUpdater = async (token: any, user: any, event: H3Event) => {
  const currentSession = await useSession(event, {
    // PASSWORD TO change as .env variable
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
  })
  await currentSession.clear()

  const newSession = await useSession(event, {
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
    cookie: {
      httpOnly: true,
      secure: true,
      sameSite: "strict",
    },
    maxAge: 60 * 60 * 24 * 7,
    generateId: () => { return token }
  })
  await newSession.update({
    userInfos: user,
    sessionToken: token
  })

  return newSession
}

export const sessionDeleter = async (event: H3Event) => {
  const currentSession = await useSession(event, {
    // PASSWORD TO change as .env variable
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
  })
  await currentSession.clear()
}

export const sessionGetter = async (event: H3Event) => {
  const sessionServer = await getSession(event, {
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
  })
  return sessionServer
}