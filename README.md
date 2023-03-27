# 목차

- [목차](#목차)
- [명령어 만들기](#명령어-만들기)
  - [Cobra 설치하기](#cobra-설치하기)
  - [CLI 프로그램 초기화하기](#cli-프로그램-초기화하기)
    - [모듈 초기화하기](#모듈-초기화하기)
    - [Cobra CLI 어플리케이션 초기화하기](#cobra-cli-어플리케이션-초기화하기)
    - [명령어 추가하기](#명령어-추가하기)
    - [플래그 추가하기](#플래그-추가하기)
- [Golang과 친해지기](#golang과-친해지기)
  - [의존성을 설치하는 방법](#의존성을-설치하는-방법)
    - [1. go 모듈 초기화하기](#1-go-모듈-초기화하기)
    - [2. 모듈(dependency) 설치하기](#2-모듈dependency-설치하기)
    - [3. 설치한 모듈을 사용하고자 하는 코드에 `import`하기](#3-설치한-모듈을-사용하고자-하는-코드에-import하기)
    - [4. `import`한 모듈을 원하는 코드 영역에서 사용하기](#4-import한-모듈을-원하는-코드-영역에서-사용하기)
    - [5. go mod tidy 명령으로 `go.mod` 파일과 `go.sum` 파일을 저장하기.](#5-go-mod-tidy-명령으로-gomod-파일과-gosum-파일을-저장하기)
  - [유용하고 강력한 기본 모듈](#유용하고-강력한-기본-모듈)
- [저장소 설정하기](#저장소-설정하기)
  - [포함된 모듈 목록](#포함된-모듈-목록)
  - [자동화](#자동화)

# 명령어 만들기

## Cobra 설치하기

go가 설치되어 있다면, go의 cli 프로그램 생성 도구인 [cobra](https://github.com/spf13/cobra#installing)를 설치 해 주세요.

```shell
go install github.com/spf13/cobra-cli@latest
```

## CLI 프로그램 초기화하기

### 모듈 초기화하기

> 이미 모듈이 있다면 이 과정은 생략 해 주세요.

1. 디렉토리 만들기
2. 만든 디렉토리로 `cd` 해주기
3. `go mod init <MODNAME>`으로 모듈 생성하기

### Cobra CLI 어플리케이션 초기화하기

```shell
cobra-cli init
```

### 명령어 추가하기

> `cobra-cli add command-name` 으로 명령어를 추가해요.
> `cobra-cli add child-command-name -p 'parent-command-nameCmd'` 으로 child 명령어를 만들 수 있어요.

### 플래그 추가하기

- [Persistent Flag](https://github.com/spf13/cobra/blob/main/user_guide.md#persistent-flags)와 [Local Flag](https://github.com/spf13/cobra/blob/main/user_guide.md#local-flags)도 추가 할 수 있어요.

```go
// Persistent 플래그는 명령어 자신 뿐만 아니라 하위 명령어에도 적용되는 플래그를 말해요.
// 글로벌 플래그를 만들고 싶다면, 루트 명령어에 Persistent 플래그를 적용 해 주면 된답니다.
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

// 원하는 명령어에만 적용 할 수 있는 플래그도 지정 가능합니다!
localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

# Golang과 친해지기

## 의존성을 설치하는 방법

### 1. go 모듈 초기화하기

```shell
go mod init <module-name>
```

### 2. 모듈(dependency) 설치하기

```shell
go get github.com/user/package
```

- 설치 된 모듈은 `go.mod` 파일에 표시 됩니다.

### 3. 설치한 모듈을 사용하고자 하는 코드에 `import`하기

```go
import {
  "github/user/package"
}
```

### 4. `import`한 모듈을 원하는 코드 영역에서 사용하기

```go
func mai() {
  // Your code here
}
```

### 5. go mod tidy 명령으로 `go.mod` 파일과 `go.sum` 파일을 저장하기.

```shell
go mod tidy
```

## 유용하고 강력한 기본 모듈

- os/exec 모듈을 임포트 한 뒤 `exec.Command(name string, arg ...string) \*cmd` [명령어](https://pkg.go.dev/os/exec#Command)로 각 운영체제에 맞는 command line command를 사용 할 수 있어요.

```go
// (...)
func main () {
  cmd := exec.Command("tr", "a-z", "A-Z")     // 소문자를 대문자로 바꾸기
}
// (...)
```

# 저장소 설정하기

## 포함된 모듈 목록

|                           모듈                           |                       용도                       |
| :------------------------------------------------------: | :----------------------------------------------: |
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) |   Functional 하고 Stateful한 터미널 앱 만들기    |
|  [Lip gloss](https://github.com/charmbracelet/lipgloss)  | 좋은 터미널 레이아웃을 구성하기 위한 스타일 정의 |

## 자동화

| 파일                             | 용도                                                                                                                                                                                                                                                                                                                                                                  |
| -------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `.cobra.yaml`                    | 여기에 제작자와 라이선스를 입력 해 두면 매번 제작자와 라이선스를 수정하지 않아도 됩니다!                                                                                                                                                                                                                                                                              |
| `.pre-commit-config.yaml`        | [pre-commit](https://pre-commit.com/) 설정 파일.<br/> `pre-commit install`로 git hook을 설치 할 수 있고, `pre-commit sample-config`로 이 파일을 생성 할 수 있습니다.                                                                                                                                                                                                  |
| `.gorelaser.yaml`                | [GoReleaser](https://goreleaser.com/) 설정 파일. `goreleaser init` 명령어로 생성합니다. <br>컴파일, 릴리즈 노트 생성, Homebrew Formulae를 생성하고 나의 Homebrew Tap 저장소에 배포하는 작업을 자동화 하는 데 사용됩니다.                                                                                                                                              |
| `.github/workflows/release.yml`. | Github Action을 이용해 [GoReleaser의 작업을 자동화](https://goreleaser.com/ci/actions/?h=github+ac) 하기 위한 설정 파일. <br/> `Homebrew-tapName`으로 [자신](https://github.com/mindulle/homebrew-mindulle)의 [brew tap](https://docs.brew.sh/Taps)을 만들어 두었다면, 이 템플릿에서 만든 명령어를 자신의 tap에 자동 배포해주는 Github Action을 위한 설정 파일이에요. |
