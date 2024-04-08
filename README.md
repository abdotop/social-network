# 👋 Social Network

## Description
This project consist of creating a facebook like web application by using any of existing js framework for the frontend and golang for the backend.
To know more about this project, please visit this [link](https://github.com/01-edu/public/tree/master/subjects/social-network)

## Features
+ **Followers**

    Users have the ability to follow each others.

+ **Profile**

    Each user has a profile presenting some information about him/her.

+ **Posts**

    Users  have the ability to post their best moments with their friends.

+ **Groups**

    Users can create or join groups, enabling them to be part of a community that shares their interests.

+ **Notification**

    Users will be notified on every event that concern him/her.

+ **Chats**

    Users can exchange messages with each others


## Requirements
+ **Nuxtjs**

    Nuxt is a free and open source JavaScript library based on Vue.js, Nitro, and Vite. [Wikipedia](https://en.wikipedia.org/wiki/Nuxt.js)

    - [documentation](https://nuxt.com/docs)

    
+ **Tailwind**

    Tailwind CSS is an open source CSS framework. The main feature of this library is that, unlike other CSS frameworks like Bootstrap, it does not provide a series of predefined classes for elements such as buttons or tables. [Wikipedia](https://en.wikipedia.org/wiki/Tailwind_CSS)

    - [documantation](https://tailwindcss.com/docs/installation/framework-guides)

+ **Golang**

    Go is a statically typed, compiled high-level programming language designed at Google. [Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language))

    - [documentation](https://go.dev/doc/)

+ **SQlite**

    SQLite is a database engine written in the C programming language. [Wikipedia](https://en.wikipedia.org/wiki/SQLite)

    - [documentation](https://www.sqlite.org/docs.html)

## Installation
Open a terminal, go to the root directory of the project and run `make install`. This will install all the dependencies.

## Usage
+ **Frontend**

To start the frontend you can follow one of the following steps:

1. go to the frontend directory `cd frontend` and run this command:
`npm run dev`

2. from the root of the project run this command:
`make start-frontend`


+ **Backend**

To start the backend server you can follow one of the following steps:

1. go to the backend directory `cd backend` and run this command:
`go run .`

2. from the root of the project run this command:
`make start-backend`

You can also run `make start` to start both, the **backend** server and **frontend** server.


+ **Database**
Run this command to create a new migration logic:

`migrate create -seq -ext=.sql -dir=./pkg/db/migrations/sqlite <table_name>`

`-seq`: to use number for ordering

`-ext=<file>`: specify the migration file extension where `<file>` is the migration file.

`-dir=<directory>`: directory to the migrations directory

## Contributors

1. [Papa Alassane Ba](https://learn.zone01dakar.sn/git/pba)
2. [Bounama Coulibaly](https://learn.zone01dakar.sn/git/bcoulibal)
3. [Cheikh Ahmadou Tidiane Diallo](https://learn.zone01dakar.sn/git/chdiallo)
4. [Serigne Saliou Mbacké Mbaye](https://learn.zone01dakar.sn/git/serignmbaye)
5. [Abdou Top](https://learn.zone01dakar.sn/git/abdotop)
6. [Mouhamed Sarr](https://learn.zone01dakar.sn/git/mohamedsarr0)
