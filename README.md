## [Get this title for $10 on Packt's Spring Sale](https://www.packt.com/B12404?utm_source=github&utm_medium=packt-github-repo&utm_campaign=spring_10_dollar_2022)
-----
For a limited period, all eBooks and Videos are only $10. All the practical content you need \- by developers, for developers

# Go Programming Cookbook - Second Edition 

<a href="https://www.packtpub.com/in/application-development/go-programming-cookbook-second-edition?utm_source=github&utm_medium=repository&utm_campaign=9781789800982"><img src="https://www.packtpub.com/media/catalog/product/cache/e4d64343b1bc593f1c5348fe05efa4a6/9/7/9781789800982-original.jpeg" alt="Go Programming Cookbook - Second Edition " height="256px" align="right"></a>

This is the code repository for [Go Programming Cookbook - Second Edition ](https://www.packtpub.com/in/application-development/go-programming-cookbook-second-edition?utm_source=github&utm_medium=repository&utm_campaign=9781789800982), published by Packt.

**Over 85 recipes to build modular, readable, and testable Golang applications across various domains**

## What is this book about?
Go (or Golang) is a statically typed programming language developed at Google. Known for its vast standard library, it also provides features such as garbage collection, type safety, dynamic-typing capabilities, and additional built-in types. This book will serve as a reference while implementing Go features to build your own applications.


This book covers the following exciting features:
* Work with third-party Go projects and modify them for your use 
* Write Go code using modern best practices 
* Manage your dependencies with the new Go module system 
* Solve common problems encountered when dealing with backend systems or DevOps 
* Explore the Go standard library and its uses 
Test, profile, and fine-tune Go applications

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1789800986) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
                b, err := ioutil.ReadAll(r)
                if err != nil {
                    return "", err
                }
                return string(b), nil
        }
```

**Following is what you need for this book:**
This book is aimed for web developers, programmers, and enterprise developers. Basic knowledge of the Go language is assumed. Experience with backend application development is not necessary, but may help understand the motivation behind some of the recipes.

This book serves as a good reference for Go developers who are already proficient but need a quick reminder, example, or reference. With the open source repository, it should be possible to share these examples quickly with a team as well. If you are looking for quick solutions to common and not-so-common problems in Go programming, this book is for you.

## Errata
The following content at Page 323 under *How it works...* should read as:
Consul provides a robust Go API library. It can feel daunting when you're starting for the first time, but this recipe shows how you might approach wrapping it. Configuring Consul further is beyond the scope of this recipe; this shows the basics of registering a service and querying for other services when given a key and tag. It would be possible to use this to register new microservices at startup time, query for all dependent services, and deregister at shutdown. You may also want to cache this information so that you're not hitting Consul for every request, but this recipe provides the basic tools that you can expand upon. The Consul agent also makes these repeated requests
fast and efficient (https://www.consul.io/intro/getting-started/agent.html).

With the following software and hardware list you can run all code files present in the book (Chapter 1-).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| All | Go - latest version | Windows, macOS and Linux |
| All | Unix Shell - any version | Windows, macOS and Linux |
| 6 | MySQL - latest version | Windows, macOS and Linux |
| 6 | Redis - latest version | Windows, macOS and Linux |
| 6 | MongoDB - latest version | Windows, macOS and Linux |
| 11 | consul.io - latest version | Windows, macOS and Linux |
| 11 | Docker - latest community edition version | Windows, macOS and Linux |
| 11 | Prometheus - latest version | Windows, macOS and Linux |

### Related products
* Learn Data Structures and Algorithms with Golang  [[Packt]](https://www.packtpub.com/application-development/learn-data-structures-and-algorithms-golang?utm_source=github&utm_medium=repository&utm_campaign=9781789618501) [[Amazon]](https://www.amazon.com/dp/1789618509)

* Machine Learning With Go  [[Packt]](https://www.packtpub.com/big-data-and-business-intelligence/machine-learning-go?utm_source=github&utm_medium=repository&utm_campaign=9781785882104) [[Amazon]](https://www.amazon.com/dp/1785882104)

## Get to Know the Author
**Aaron Torres**
received his master's degree in computer science from the New Mexico Institute of Mining and Technology. He has worked on distributed systems in high-performance computing and in large-scale web and microservices applications. He currently leads a team of Go developers that refines and focuses on Go best practices with an emphasis on continuous delivery and automated testing.

Aaron has published a number of papers and has several patents in the area of storage and I/O. He is passionate about sharing his knowledge and ideas with others. He is also a huge fan of the Go language and open source for backend systems and development.

### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.


