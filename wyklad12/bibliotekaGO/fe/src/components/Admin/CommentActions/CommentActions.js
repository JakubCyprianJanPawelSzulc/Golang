import './CommentActions.scss';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { useEffect, useRef, useState } from 'react';
import { booksService } from '../../../services/books.service';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { Dialog } from 'primereact/dialog';
import { InputTextarea } from 'primereact/inputtextarea';
import { useForm, Controller } from 'react-hook-form';
import { Dropdown } from 'primereact/dropdown';
import { userService } from '../../../services/user.service';

export function CommentActions() {
  const [comments, setComments] = useState(null);
  const toast = useRef(null);
  const [showAddDialog, setShowAddDialog] = useState(false);

  const [editMode, setEditMode] = useState({ mode: false, comment: null });

  const [users, setUsers] = useState(null);
  const [books, setBooks] = useState(null);
  const [selectedUser, setSelectedUser] = useState(null);
  const [selectedBook, setSelectedBook] = useState(null);

  const defaultValues = {
    comment: null
  };

  const {
    control,
    register,
    handleSubmit,
    setError,
    getValues,
    setValue,
    reset,
    formState: { errors }
  } = useForm();

  const resetForm = () => {
    reset(defaultValues);
    setSelectedUser(null);
    setSelectedBook(null);
  };

  const getComments = () => {
    booksService.getAllComments().then((e) => {
      setComments(e);
    });
  };

  useEffect(() => {
    getComments();

    userService.getAllUsers().then((e) => {
      const withFullName = e.map((el) => {
        return { ...el, fullName: el.name + ' ' + el.surName };
      });
      setUsers(withFullName);
    });

    booksService.getAllBooks().then((e) => {
      setBooks(e.data);
      console.log(e.data);
    });
  }, []);

  useEffect(() => {
    if (editMode.mode) {
      reset(editMode.comment);
      setShowAddDialog(true);
    }
  }, [editMode]);

  const buttonTemplate = (data) => (
    <div className="buttons">
      <Button className="p-button-rounded p-button-text" onClick={() => toggleEditMode(data)} icon="pi pi-pencil" />
      <Button className="p-button-rounded p-button-text" onClick={() => deleteComment(data)} icon="pi pi-times" />
    </div>
  );

  const deleteComment = (data) => {
    booksService.deleteCommentbyId(data._id).then((e) => {
      getComments();
      toast.current.show({
        severity: 'success',
        summary: 'Usunięto',
        detail: 'Pomyślnie usunięto komentarz.',
        life: 3000
      });
    });
  };

  const getFullName = (rowData) => {
    return rowData.userName + ' ' + rowData.userSurName;
  };

  const addCommentFooter = (name) => {
    return (
      <div>
        <Button
          label="Anuluj"
          onClick={() => {
            setShowAddDialog(false);
            setEditMode({ mode: false, book: null });
            resetForm();
          }}
          className="p-button-text"
        />
        <Button label="Zapisz" form="add-form" autoFocus />
      </div>
    );
  };

  const addCommentSubmit = (data) => {
    if (selectedBook && selectedUser) {
      const toInsert = {
        userId: selectedUser.userId,
        bookId: selectedBook._id,
        comment: data.comment,
        userName: selectedUser.name,
        userSurName: selectedUser.surName
      };
      booksService.addCommentByAdmin(toInsert).then((e) => {
        getComments();
        resetForm();
        setShowAddDialog(false);
        toast.current.show({
          severity: 'success',
          summary: 'Dodano',
          detail: 'Pomyślnie dodano komentarz.',
          life: 3000
        });
      });
    }
  };

  const editCommentSubmit = (data) => {
    booksService.updateComment(data).then((e) => {
      getComments();
      resetForm();
      setShowAddDialog(false);
      setEditMode({ mode: false, comment: null });
      toast.current.show({
        severity: 'success',
        summary: 'Zaktualizowano',
        detail: 'Pomyślnie zaktualizowano komentarz.',
        life: 3000
      });
    });
  };

  const toggleEditMode = (passedComment) => {
    setShowAddDialog();
    setEditMode({ mode: true, comment: passedComment });
  };

  return (
    <div>
      <Dialog
        closable={false}
        header="Komentarz"
        footer={addCommentFooter}
        visible={showAddDialog}
        style={{ width: '50vw' }}
        onHide={() => setShowAddDialog(false)}
      >
        {!editMode.mode && (
          <div className="add-book-drop">
            <Dropdown
              value={selectedUser}
              options={users}
              onChange={(e) => setSelectedUser(e.value)}
              optionLabel="fullName"
              placeholder="Wybierz użytkownika"
            />

            <Dropdown
              value={selectedBook}
              options={books}
              onChange={(e) => setSelectedBook(e.value)}
              optionLabel="title"
              placeholder="Wybierz Książkę"
            />
          </div>
        )}

        <form id="add-form" onSubmit={handleSubmit(editMode.mode ? editCommentSubmit : addCommentSubmit)}>
          <div className="login-box">
            <div className="flex-col-gap">
              <InputTextarea
                rows={5}
                cols={30}
                autoResize
                id="description"
                {...register('comment', {
                  required: true
                })}
                className={errors.comment && 'p-invalid'}
                placeholder="Opis"
              />

              {errors.comment?.type === 'required' && (
                <small id="comment-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              {errors['quantity'] && (
                <small id="quantity-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
            </div>
          </div>
        </form>
      </Dialog>

      <div className="card card-comment">
        <Button
          label="Dodaj Komentarz"
          onClick={() => setShowAddDialog(true)}
          className="p-button-rounded p-button-text"
        />
        <Toast ref={toast} />
        <DataTable value={comments} responsiveLayout="scroll">
          <Column field="bookTitle" header="Tytuł"></Column>
          <Column field="comment" header="Komentarz"></Column>
          <Column header="Użytkownik" body={getFullName}></Column>

          <Column header="Akcje" body={buttonTemplate}></Column>
        </DataTable>
      </div>
    </div>
  );
}
