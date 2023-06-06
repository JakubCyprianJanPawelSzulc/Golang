import './UserActions.scss';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { useEffect, useRef, useState } from 'react';
import { booksService } from '../../../services/books.service';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { InputText } from 'primereact/inputtext';
import { Dialog } from 'primereact/dialog';
import { InputTextarea } from 'primereact/inputtextarea';
import { useForm, Controller } from 'react-hook-form';
import { Dropdown } from 'primereact/dropdown';
import { userService } from '../../../services/user.service';

export function UserActions() {
  const [users, setUsers] = useState(null);
  const toast = useRef(null);
  const [showAddDialog, setShowAddDialog] = useState(false);

  const [editMode, setEditMode] = useState({ mode: false, user: null });

  const defaultValues = {
    name: null,
    surName: null
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
  };

  const getUsers = () => {
    userService.getAllUsers().then((e) => {
      setUsers(e);
    });
  };

  useEffect(() => {
    getUsers();
  }, []);

  useEffect(() => {
    if (editMode.mode) {
      reset(editMode.user);
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
    userService.deleteUserbyId(data._id).then((e) => {
      getUsers();
      toast.current.show({
        severity: 'success',
        summary: 'Usunięto',
        detail: 'Pomyślnie usunięto użytkownika.',
        life: 3000
      });
    });
  };

  const addCommentFooter = (name) => {
    return (
      <div>
        <Button
          label="Anuluj"
          onClick={() => {
            setShowAddDialog(false);
            setEditMode({ mode: false, user: null });
            resetForm();
          }}
          className="p-button-text"
        />
        <Button label="Zapisz" form="add-form" autoFocus />
      </div>
    );
  };

  const addCommentSubmit = (data) => {};

  const editCommentSubmit = (data) => {
    userService.editUserNameAndSurname({ userId: data._id, name: data.name, surName: data.surName }).then((e) => {
      getUsers();
      setShowAddDialog(false);
      reset();
      toast.current.show({
        severity: 'success',
        summary: 'Edytowano',
        detail: 'Pomyślnie edytowano użytkownika.',
        life: 3000
      });
    });
  };

  const toggleEditMode = (passedUser) => {
    setShowAddDialog();
    setEditMode({ mode: true, user: passedUser });
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
        <form id="add-form" onSubmit={handleSubmit(editCommentSubmit)}>
          <div className="login-box">
            <div className="flex-col-gap">
              <InputText
                id="name"
                {...register('name', {
                  required: true
                })}
                className={errors.name && 'p-invalid'}
                placeholder="Imię"
              />

              {errors.name?.type === 'required' && (
                <small id="name-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <InputText
                id="surName"
                {...register('surName', {
                  required: true
                })}
                className={errors.surName && 'p-invalid'}
                placeholder="Nazwisko"
              />

              {errors.surName?.type === 'required' && (
                <small id="surName-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
            </div>
          </div>
        </form>
      </Dialog>

      <div className="card card-comment">
        <Toast ref={toast} />
        <DataTable value={users} responsiveLayout="scroll">
          <Column field="name" header="Imie"></Column>
          <Column field="surName" header="Nazwisko"></Column>
          <Column field="email" header="e-mail"></Column>
          <Column field="postal" header="Kod pocztowy"></Column>
          <Column field="city" header="Miasto"></Column>
          <Column header="Akcje" body={buttonTemplate}></Column>
        </DataTable>
      </div>
    </div>
  );
}
