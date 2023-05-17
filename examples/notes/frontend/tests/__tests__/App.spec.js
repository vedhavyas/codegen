import { render, fireEvent } from '@testing-library/svelte';
import App from '../../src/App.svelte';
import * as noteAPI from '../../src/components/noteAPI.js';

jest.mock('../../src/components/noteAPI.js');

describe('App component', () => {
  test('renders notes list and add note button', async () => {
    const { getByText, getByRole } = render(App);

    expect(getByText('Notes List')).toBeInTheDocument();
    const addNoteButton = getByRole('button', { name: 'Add Note' });
    expect(addNoteButton).toBeInTheDocument();

    noteAPI.getNotes.mockResolvedValue([]);

    await fireEvent.click(addNoteButton);

    expect(noteAPI.addNote).toHaveBeenCalled();
  });

  test('edits a note properly', async () => {
    const { getByTestId } = render(App);

    const editNoteComponent = getByTestId('edit-note');
    expect(editNoteComponent).toBeInTheDocument();

    fireEvent.submit(editNoteComponent);

    expect(noteAPI.updateNote).toHaveBeenCalled();
  });

  test('deletes a note properly', async () => {
    const { getByTestId } = render(App);

    const deleteNoteComponent = getByTestId('delete-note');
    expect(deleteNoteComponent).toBeInTheDocument();

    fireEvent.submit(deleteNoteComponent);

    expect(noteAPI.deleteNote).toHaveBeenCalled();
  });
});