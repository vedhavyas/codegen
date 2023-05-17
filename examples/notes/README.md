├─ backend
│   ├── Cargo.lock
│   ├── Cargo.toml
│   ├── src
│   │   ├── main.rs
│   │   ├── db.rs
│   │   └── models.rs
│   └── tests
│       └── test_mod.rs
├─ frontend
│   ├── package.json
│   ├── src
│   │   ├── App.svelte
│   │   ├── Note.svelte
│   │   ├── components
│   │   │   ├── AddNote.svelte
│   │   │   ├── DeleteNote.svelte
│   │   │   ├── EditNote.svelte
│   │   │   ├── NotesList.svelte
│   │   │   └── noteAPI.js
│   └── tests
│       └── __tests__
│           └── App.spec.js
├─ infra
│   ├── docker-compose.dev.yml
│   └── docker-compose.prod.yml
├─ .github
│   └── workflows
│       └── ci-cd.yml
└─ README.md
```

## Setup

### Backend
To setup and run the backend, use the following commands:

```bash
cd backend
cargo install --path .
cargo run
```

### Frontend
To setup and run the frontend, use the following commands:

```bash
cd frontend
npm install
npm run dev
```

## Testing

### Backend
To run the backend tests, use the following commands:

```bash
cd backend
cargo test
```

### Frontend
To run the frontend tests, use the following commands:

```bash
cd frontend
npm test
