Resources:

- https://www.youtube.com/watch?v=MzTcsI6tn-0
- https://dave.cheney.net/practical-go/presentations/qcon-china.html#_a_good_package_starts_with_its_name

Not specified:

- Who can create rooms? I guess artists? or admins?
- Does the user joins a room? or is being added by someone?

I'm gonna assume all this call need to be authenticated. Just by header.

Endpoints:

[x] Create a room
[x] Add user to room
[x] Send a message to a room
[] Delete a message from a room
[] Get latest messages from a room
[] Get information about a given room

Requisites:

Create a service to expose an API.
Automated tests should be added to verify the correct behavior of the API.
Use a database of your choice (or no database at all if you prefer).
Bonus: Expose documentation.
Bonus: Dockerize your application.
