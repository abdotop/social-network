interface UserInfos {
    id: String,
    email: String,
    password: String,
    firstName: String,
    lastName: String,
    dateOfBirth: String,
    avatarImage: String,
    nickname: String,
    aboutMe: String,
    isPublic: true,
    createdAt: String,
    updatedAt: String,
    deletedAt: { Time: '0001-01-01T00:00:00Z', Valid: false }
}


id: "0c1a0f8c-98f4-448b-a2dd-2f33c6bcec48"
password: "$2a$10$VqY6ytlW.9.qx40BZNa5uuTtaw7Anse0kxoBUtoTboA36b8sysm1O"

export const secure = (user: UserInfos) => {
    return {
        email: user.email,
        firstName: user.firstName, 
        lastName: user.lastName,
        dateOfBirth: user.dateOfBirth.split("T00:00:00Z")[0],
        avatarImage: user.avatarImage, 
        nickname: user.nickname, 
        aboutMe: user.aboutMe, 
        isPublic: user.isPublic, 
        createdAt: user.createdAt, 
        updatedAt: user.updatedAt, 
    }
}

export const encoder = (token: string) => {
    return btoa(token)
}

export const decoder = (token: string) => {
    return atob(token)
}