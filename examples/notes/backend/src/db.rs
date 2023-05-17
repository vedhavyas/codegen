use rusqlite::{Connection, Result, params};
use crate::models::Note;

pub fn run_db() -> Result<Connection> {
    let conn = Connection::open("notes.db")?;
    conn.execute(
        "CREATE TABLE IF NOT EXISTS notes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )",
        [],
    )?;

    Ok(conn)
}

pub fn create_note(conn: &Connection, note: &Note) -> Result<usize> {
    conn.execute(
        "INSERT INTO notes (title, content) VALUES (?1, ?2)",
        params![note.title, note.content],
    )
}

pub fn update_note(conn: &Connection, note: &Note) -> Result<usize> {
    conn.execute(
        "UPDATE notes SET title = ?1, content = ?2 WHERE id = ?3",
        params![note.title, note.content, note.id],
    )
}

pub fn delete_note(conn: &Connection, id: i32) -> Result<usize> {
    conn.execute("DELETE FROM notes WHERE id = ?1", params![id])
}

pub fn list_notes(conn: &Connection) -> Result<Vec<Note>> {
    let mut stmt = conn.prepare("SELECT id, title, content, created_at FROM notes")?;
    let note_iter = stmt.query_map([], |row| {
        Ok(Note {
            id: row.get(0)?,
            title: row.get(1)?,
            content: row.get(2)?,
            created_at: row.get(3)?,
        })
    })?;

    let mut notes = Vec::new();
    for note in note_iter {
        notes.push(note?);
    }

    Ok(notes)
}