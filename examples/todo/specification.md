A simple Todo List application that allows users to add, edit, mark as complete, and delete tasks. The application is developed using a Svelete as frontend and Go as backend to store the state.
Specifications:
1. Tasks should be displayed in a list view
2. Option to add a new task
3. Option to edit a task
4. Option to mark tasks as complete/incomplete
5. Option to delete tasks
6. Include unit tests for components and logic
7. Use modern CSS components for cleaner UI

Project structure:
- backend
- frontend
- infra
- ...

Backend:
- Use go 1.19
- Use `github.com/vedhavyas/todo` as the go module name
- Keep all the components under `backend` folder without any inner packages.
- Use gorilla mux with standard http router. Do not use Gin or similar frameworks.
- Use Sqlite to store data.

CI/CD:
- Use github actions
- Use docker compose with production and dev variants

Note:
- Implement all the necessary components.
- The app should be production ready.
- Do not give any unimplemented or mocked logic.