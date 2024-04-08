export interface User {
  id: string;
  isLoggedIn: boolean;
  avatarImage?: string;
  isPublic: boolean;
  firstName: string;
  lastName: string;
  nickname: string;
  email: string;
  password: string;
  dateOfBirth: string;
  aboutMe: string;
  // pseudo: string;
}

export interface ServerResponse<T> {
  status: string;
  session?: string;
  message: string;
  data?: T;
}

export interface Group {
  ID: string,
  Title: string,
  Description: string,
  BannerURL: string,
  CreatorID: string,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string,
  GroupMembers: GroupMember[]
}

export interface GroupMember {
  ID: string,
  GroupID: string,
  MemberID: string,
  Status: string,
  Role: string,
  User: UserWithoutPassword,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string
}

export interface GroupMessage {
  ID: string,
  GroupID: string,
  SenderID: string,
  Sender: UserWithoutPassword,
  Content: string,
  CreatedAt: string,
  UpdatedAt: string,
  DeletedAt: string
}

export interface Event {
  ID: string,
  GroupID: string,
  title: string,
  description: string,
  CreatorID: string,
  date_time: string,
  Participants: EventParticipant[]

}

export interface EventParticipant {
  ID: string,
  MemberID: string,
  EventID: string,
  response: string,
  User: UserWithoutPassword
}

export interface Post {
  id: string,
  group_id: string,
  user_id: string,
  title: string,
  content: string,
  image_url: string,
  privacy: string,
  followersSelectedID: string,
  created_at: string,
  updated_at: string,
  deleted_at: Date,
  commets: Comment[]
}

export interface Comment {
  content: string,
  createdAt: string,
  id: string,
  imageUrl: string,
  post_id: string,
  userAvatarImageUrl: string,
  userCompleteName: string,
  userOwnerNickname: string
}

export interface Invitation {
  Group: Group,
  MemberId: string
}

export type UserWithoutPassword = Omit<User, "password">;