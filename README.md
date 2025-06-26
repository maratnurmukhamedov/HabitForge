# HabitForge
**HabitForge** is a simple habit tracker written in Go (and also my first project), supporting two modes:
 
- **CLI (Command-Line Interface)** — local interactive usage
- **HTTP API server** — a lightweight REST API ready for integration

All habit data is stored in a local file: `habits.json`.

---

## Getting Started

Make sure you're inside the project directory.

### Run CLI mode

```bash
go run . cli
```

### Run HTTP server mode

```bash
go run . server
```

---

## CLI Menu

```
1. Add new habit
2. Describe habit progress
3. See all habits
4. Check your streak
5. Delete a habit
0. Quit
```

- Data is stored in `habits.json`
- CLI and server share the same logic and storage

---

## REST API

Server runs at: `http://localhost:8080`

### Available Endpoints

| Method | Endpoint             | Description                      |
|--------|----------------------|----------------------------------|
| GET    | `/habits`            | Get all habits                   |
| POST   | `/habits`            | Add a new habit                  |
| DELETE | `/habits?name=X`     | Delete a habit by name           |
| GET    | `/streaks`           | Get longest streaks per habit    |
| POST   | `/shutdown`          | Gracefully shut down the server  |

---

## API Examples (PowerShell with curl.exe - working on Windows 10+)

### Add a habit

```powershell
curl.exe -X POST http://localhost:8080/habits -H "Content-Type: application/json" -d '{\"Name\":\"Workout\",\"Priority\":2}'
```

### Get all habits

```powershell
curl.exe http://localhost:8080/habits
```

### Delete a habit

```powershell
curl.exe -X DELETE "http://localhost:8080/habits?name=Workout"
```

### Get streaks

```powershell
curl.exe http://localhost:8080/streaks
```

### Shutdown the server

```powershell
curl.exe -X POST http://localhost:8080/shutdown
```

---

## Streak Logic

A streak is calculated as the **longest number of consecutive days** a habit was marked as completed.

Example output:
```
Drink water: 2 days
Workout: 0 days
```


