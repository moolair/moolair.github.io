# My Blog

Welcome to my blog! Visit the [main page](https://moolair.github.io/index.html) to see the posts.

## About
This is a personal blog built with Firebase and GitHub Pages.

##WIP 현재 프로젝트 구조
project-name/
├── cmd/
│   ├── blog-server/
│   │   ├── main.go         # Firebase와 API 서버 초기화
├── internal/
│   ├── blog/
│   │   ├── handler.go      # /posts 핸들러 구현
├── pkg/
│   ├── db/
│   │   ├── firebase.go     # Firebase 초기화 및 설정
├── web/
│   ├── templates/
│   ├── static/
└── go.mod                  # Go 모듈 설정
