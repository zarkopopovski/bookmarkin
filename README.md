# BookmarkIn

## BookmarkIn - personal bookmark api

**BookmarkIn** is basic personal bookmark service based on SQLite RDBMS as backend storage.

### WHY?
Because i am tired of synchronizing bookmarks between browsers and cloud services so I decide to write my own open bookmark manager that will cover my needs and in the same time can be configured to serve multiple users for the people with who i want to share them.

### WHY SQLite and not MySQL / MariaDB / PostgreSQL?
The main idea behind this service is to be used on your own cloud as your personal bookmark service which can be also configured to serve multiple users. On other side SQLite RDBMS is prooven database engine which is wide spread, its integrated almost everywhere, in every piece of software and device but also is portable and easy for backup which mean for this personal service SQLite is more then adequate.

### WHY API?
**BookmarkIn** is implemented as Restful API so its up to you what frontend you will use or in which language you will implement :).