# github-activity

https://roadmap.sh/projects/github-user-activity

---

## Set your GitHub token to access private repos (Optional)

```sh
export GITHUB_TOKEN=your_github_token
```

Or set `GITHUB_TOKEN` in `.env` file.

## Run

```sh
go run main.go <username>
```

## Build

```sh
go build
```

## Example Commands

### List users events

```sh
./github-activity JakeWharton
```

```sh
./github-activity frenck
```

```sh
./github-activity cheshire137
```

```sh
./github-activity gaearon
```
