## REQUIREMENTS
1. **Followers:**
   - Users should be able to follow and unfollow each other.
   - There should be a mechanism for sending follow requests and accepting or declining them.
   - Data structures needed:
     - Users table
     - Followers table (to store follower-followee relationships)
     - Follow requests table (to store pending follow requests)

2. **Profile:**
   - User profiles should contain information such as email, first name, last name, date of birth, avatar/image, nickname, and about me.
   - Profiles can be public or private, with different visibility settings.
   - Data structures needed:
     - Users table (to store user profile information)
     - Profile privacy settings table (to store privacy settings for each user)

3. **Posts:**
   - Users should be able to create, retrieve, update, and delete posts.
   - Posts can have privacy settings (public, private, almost private) and may include images or GIFs.
   - Data structures needed:
     - Posts table
     - Post privacy settings table

4. **Groups:**
   - Users should be able to create groups and invite others to join.
   - Groups can have titles, descriptions, and members.
   - Users can create events within groups.
   - Data structures needed:
     - Groups table
     - Group members table
     - Events table

5. **Chats:**
   - Users should be able to send private messages to each other.
   - Users can also participate in group chats.
   - Real-time communication using WebSockets is required.
   - Data structures needed:
     - Private messages table
     - Group messages table
     - Chats table (to store chat participants)

6. **Notifications:**
   - Users should receive notifications for various events, such as:
    - follow requests
    - group invitations
    - new messages.
   - Data structures needed:
     - Notifications table

7. **Authentication:**
   - Users should be able to register, login, and logout.
   - Authentication should be handled securely using sessions and cookies.
   - Data structures needed:
     - Sessions table

8. **Database:**
   - SQLite or another database system should be used to store user data, posts, messages, etc.
   - Data structures needed:
     - Tables for users, posts, messages, notifications, etc.

9. **Security:**
   - Security measures such as:
    - password hashing,
    - input validation
    - protection against common vulnerabilities (e.g., SQL injection, XSS) are required.
   - Data structures needed:
     - User authentication tokens table

10. **Performance and Scalability:**
    - The database schema and backend architecture should be designed to handle a large number of users, posts, messages, etc., efficiently.
    - Data structures needed:
      - Proper indexing and optimization of database queries

## DATA STRUCTURE
1. **Users Table:**
   - Columns:
     - id (Primary Key)
     - email (Unique)
     - password (Hashed)
     - first_name
     - last_name
     - date_of_birth
     - avatar_image
     - nickname
     - about_me
     - is_public
     - created_at
     - updated_at
     - delete_at

2. **Followers Table:**
   - Columns:
     - id (Primary Key)
     - follower_id (Foreign Key referencing Users table)
     - followee_id (Foreign Key referencing Users table)
     - status (Enum: requested, accepted, declined)
     - created_at
     - updated_at

3. **Posts Table:**
   - Columns:
     - id (Primary Key)
     - user_id (Foreign Key referencing Users table)
     - title
     - content
     - image_url
     - privacy (Enum: public, private, almost private, unlisted)
     - created_at
     - updated_at
     - deleted_at

4. **Groups Table:**
   - Columns:
     - id (Primary Key)
     - title
     - description
     - banner_url
     - creator_id (Foreign Key referencing Users table)
     - created_at
     - updated_at

5. **Group Members Table:**
   - Columns:
     - id (Primary Key)
     - group_id (Foreign Key referencing Groups table)
     - member_id (Foreign Key referencing Users table)
     - status (Enum: invited, requesting, accepted, decline)
     - created_at
     - updated_at

6. **Group Posts Table:**
   - Columns:
     - id (Primary Key)
     - group_id (Foreign Key referencing Groups table)
     - post_id (Foreign Key referencing Posts table)
     - created_at
     - updated_at

6. **Events Table:**
   - Columns:
     - id (Primary Key)
     - group_id (Foreign Key referencing Groups table)
     - title
     - description
     - datetime
     - created_at
     - updated_at

6. **Events Participant Table:**
   - Columns:
     - id (Primary Key)
     - group_id (Foreign Key referencing Groups table)
     - member_id (Foreign Key referencing Users table)
     - response (Enum: going, not_going)
     - created_at
     - updated_at

7. **Private Messages Table:**
   - Columns:
     - id (Primary Key)
     - sender_id (Foreign Key referencing Users table)
     - receiver_id (Foreign Key referencing Users table)
     - content
     - created_at

8. **Group Messages Table:**
   - Columns:
     - id (Primary Key)
     - group_id (Foreign Key referencing Groups table)
     - sender_id (Foreign Key referencing Users table)
     - content
     - created_at

10. **Notifications Table:**
    - Columns:
      - id (Primary Key)
      - user_id (Foreign Key referencing Users table)
      - type (Enum: follow_request, group_invitation, new_message, new_event etc.)
      - message
      - created_at

11. **Sessions Table:**
    - Columns:
      - id (Primary Key)
      - user_id (Foreign Key referencing Users table)
      - session_token
      - created_at
      - expires_at
      - deleted_at

## FILE STRUCTURE
```
backend
│
├── app
├── pkg
│   ├── db
│   │   ├── migrations
│   │   │   └── sqlite
│   │   │       └── ...
│   │   └── sqlite
│   │       └── sqlite.go
│   ├── handlers
│   │   ├── authentication.go
│   │   ├── followers.go
│   │   ├── posts.go
│   │   ├── groups.go
│   │   ├── events.go
│   │   ├── messages.go
│   │   └── notifications.go
│   ├── models
│   │   ├── user.go
│   │   ├── session.go
│   │   ├── follower.go
│   │   ├── post.go
│   │   ├── group.go
│   │   ├── event.go
│   │   ├── message.go
│   │   ├── chat.go
│   │   └── notification.go
│   └── middleware
│       ├── authentication.go
│       ├── authorization.go
│       └── logging.go
├── main.go
└── go.mod
```

/post?field[title,description]
## 

##