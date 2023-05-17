import axios from "axios";

const apiBaseUrl = "http://localhost:8000";

export async function getNotes() {
  const response = await axios.get(`${apiBaseUrl}/notes`);
  return response.data;
}

export async function addNote(noteText) {
  const response = await axios.post(`${apiBaseUrl}/notes`, { text: noteText });
  return response.data;
}

export async function updateNote(noteId, updatedNoteText) {
  const response = await axios.put(`${apiBaseUrl}/notes/${noteId}`, { text: updatedNoteText });
  return response.data;
}

export async function deleteNote(noteId) {
  const response = await axios.delete(`${apiBaseUrl}/notes/${noteId}`);
  return response.data;
}