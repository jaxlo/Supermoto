# go-supermoto-tutorial
A multi page tutorial for building fullstack websites in Go

This tutorial will guide you through programing and deploying a personal website. This website will have an about section, comment wall and posts

[ ] <- Hey look, it's your own little website

There is already a go mega tutorial. And Go Supermoto has better SEO

## Chapters

Chapter 1: Go hello worlds
Chapter 2: HTML, CSS, staticfiles
Chapter 3: Templates
Chapter 4: Forms
Chapter 4: Databases (Making a comment section/wall)

-- TBD --
Chapter 4: Middleware and authentication
Chapter 5: Templ, HTMX, hot reload?
Chapter 6: Databases with SQLC
Chapter 7: Databases migrations with Goose (Combine with SQLC?)
Chapter 8: Auth/login
Chapter 9: Logging and error handling
Chapter 10: Deployment to prod (fly.io/Docker and VPS)

Later:
CSS Frameworks and/or modern CSS?
REST API for things like Flutter and React?
API services like email/sms
Include how to use Git somewhere? Maybe not.





## Future
Use sqlc and goose? https://pressly.github.io/goose/blog/2024/goose-sqlc/ (How do bools work with sqlite?)
This is a Go way of doing thigs. But with tools to speed things up

Change to the Go Fullstack Tutorial? Or include some API things?
Write a corresponding blog post with design decisions for the tutorial. (This is also a great place to get feedback)

Explore these libraries?
https://go.dev/solutions/webdev


## Jackson's thoughts
My goal with this project is to document a productive development stack in Go. (For my future projects and to help others with their projects)
When I started out, I thought the route to go would be adding libraries as needed like templ, sqlc and cookie handlers. But after evaluating the template libraries, I'm dissadisfied with what they have to offer. After evaluating and trying these third party libraries out, I came to the conclusion that the standard lib is better for this objective. (Database libraries maaaaaaay be different)

A lot of the go standard library is seen as difficult and limited. While you do have to do a little more work in Go to acomplish the same thing, it's often more elegent when I really figure out how to make it work. So I think the problem with the go standard library is really a documentation problem.

What if that's what I build? A book for go web developers (Web and fullstack). I like Go better than Python/Flask in terms of toos. But using go takes up more developement time because of difference in documentation styles. If I want to learn about go html templeates, I have to reference html/tmpl. But to learn the features of htm/tmpl, I have to reference text/tmpl. And once I'm here, the documentation is made for people who know what they are already looking for.

There is also the nice side befefit that Go code does not change much from year to year. So a program you wrote 4 years ago is very likely to complile and run well today. (That is not the case with Node/React)

TODO Rename because there is an older tutorial with this title
