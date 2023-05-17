use crate::db;
use crate::models::Note;

#[cfg(test)]
mod tests {
    use super::*;
    use std::sync::Once;

    static DB_INIT: Once = Once::new();

    fn init_database() {
        DB_INIT.call_once(|| {
            let conn = db::run_db().unwrap();
            conn.test_exclusive().unwrap();
        });
    }

    #[test]
    fn test_create_note() {
        init_database();
        let mut note = Note::new("Test note".into());
        assert!(db::create_note(&mut note).is_ok());
        assert_ne!(note.id, 0);
    }

    #[test]
    fn test_update_note() {
        init_database();
        let mut note = Note::new("Test note for update".into());
        db::create_note(&mut note).unwrap();

        note.title = "Updated test note".into();
        assert!(db::update_note(&note).is_ok());
    }

    #[test]
    fn test_delete_note() {
        init_database();
        let mut note = Note::new("Test note for delete".into());
        db::create_note(&mut note).unwrap();

        assert!(db::delete_note(note.id).is_ok());
    }

    #[test]
    fn test_list_notes() {
        init_database();
        let mut note = Note::new("Test note for listing".into());
        db::create_note(&mut note).unwrap();

        let notes = db::list_notes().unwrap();
        assert!(!notes.is_empty());
    }
}
