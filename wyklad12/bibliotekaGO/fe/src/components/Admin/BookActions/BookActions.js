import React, { useState, useEffect, useRef } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { InputText } from 'primereact/inputtext';
import { InputNumber } from 'primereact/inputnumber';
import { Dropdown } from 'primereact/dropdown';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { classNames } from 'primereact/utils';
import { Chips } from 'primereact/chips';
import './BookActions.scss';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';
import { Calendar } from 'primereact/calendar';
import { InputTextarea } from 'primereact/inputtextarea';
import { useForm, Controller } from 'react-hook-form';
import { Dialog } from 'primereact/dialog';
import { booksService } from '../../../services/books.service';

export function BookActions() {
  const [addBook, setAddBook] = useState(false);
  const [books, setBooks] = useState([]);
  const [date, setDate] = useState(null);
  const toast = useRef(null);

  const defaultValues = {
    title: null,
    description: null,
    dateRelease: null,
    genres: null,
    authors: null,
    publisher: null,
    quantity: null
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
    setValues([]);
    setAuthorsValues([]);
    setDate(null);
    setValueAmount(null);
  };

  const acceptRemove = (data) => {
    booksService.deleteBookbyId(data._id).then((e) => {
      setAddBook(false);
      toast.current.show({
        severity: 'success',
        summary: 'Usunięto',
        detail: 'Pomyślnie usunięto książkę.',
        life: 3000
      });

      getAllBooks();
    });
  };

  const rejectRemove = () => {
    setAddBook(false);
  };

  const confirmDelete = (data) => {
    confirmDialog({
      message: 'Jesteś pewnien, że chcesz usunąć tę książkę?',
      header: 'Potwierdzenie usunięcia',
      icon: 'pi pi-exclamation-triangle',
      acceptClassName: 'p-button-danger',
      accept: () => acceptRemove(data),
      reject: () => rejectRemove
    });
  };

  const [values, setValues] = useState([]);
  const [valuesAuthors, setAuthorsValues] = useState([]);
  const [valueAmount, setValueAmount] = useState(null);
  const [editMode, setEditMode] = useState({ mode: false, book: null });

  useEffect(() => {
    setValue('genres', values, { shouldValidate: true });
    setValue('authors', valuesAuthors, { shouldValidate: true });
    setValue('quantity', valueAmount, { shouldValidate: true });
    console.log(valueAmount);
  }, [values, valuesAuthors, valueAmount]);

  const getAllBooks = () => {
    booksService.getAllBooks().then((e) => {
      setBooks(e.data);
    });
  };

  useEffect(() => {
    getAllBooks();
  }, []);

  useEffect(() => {
    if (editMode.mode) {
      console.log(control);
      reset(editMode.book);
      setAddBook(true);
    }
  }, [editMode]);

  const addBookFooter = (name) => {
    return (
      <div>
        <Button
          label="Anuluj"
          onClick={() => {
            setAddBook(false);
            setEditMode({ mode: false, book: null });
            resetForm();
          }}
          className="p-button-text"
        />
        <Button label="Zapisz" form="add-form" autoFocus />
      </div>
    );
  };

  const addBookSubmit = (data) => {
    console.log(data);
    const NewData = { ...data, dateRelease: new Date(data.dateRelease).getTime() };
    booksService.addBook(NewData).then((e) => {
      setAddBook(false);
      resetForm();
      toast.current.show({ severity: 'success', summary: 'Sukces', detail: 'Pomyślnie dodano książkę', life: 3000 });
      getAllBooks();
    });
  };

  const editBookSumit = (data) => {
    const NewData = { ...data, dateRelease: new Date(data.dateRelease).getTime() };
    booksService.updateBook(NewData).then((e) => {
      console.log(e);
      if (e?.status === 201) {
        toast.current.show({
          severity: 'success',
          summary: 'Sukces',
          detail: 'Pomyślnie edytowano książkę',
          life: 3000
        });
      } else {
        toast.current.show({ severity: 'error', summary: 'Błąd', detail: 'Dany rekord nie istnieje', life: 3000 });
      }

      setAddBook(false);
      resetForm();

      setEditMode({ mode: false, book: null });
      getAllBooks();
    });
  };

  const toggleEditMode = (passedBook) => {
    setValues(passedBook.genres);
    setAuthorsValues(passedBook.authors);
    setValueAmount(passedBook.quantity);

    const newDate = new Date(parseInt(passedBook.dateRelease));

    setValue('dateRelease', newDate, { shouldValidate: true });

    setDate(newDate);

    setEditMode({ mode: true, book: passedBook });
  };

  const buttonTemplate = (data) => (
    <div className="buttons">
      <Button className="p-button-rounded p-button-text" onClick={() => toggleEditMode(data)} icon="pi pi-pencil" />
      <Button className="p-button-rounded p-button-text" onClick={() => confirmDelete(data)} icon="pi pi-times" />
    </div>
  );

  const joinArray = (data) => {
    return data.join(', ');
  };

  return (
    <div className="table-add-book">
      <Toast ref={toast} />
      <ConfirmDialog />
      <Dialog
        closable={false}
        header="Header"
        footer={addBookFooter}
        visible={addBook}
        style={{ width: '50vw' }}
        onHide={() => setAddBook(false)}
      >
        <form id="add-form" onSubmit={handleSubmit(editMode.mode ? editBookSumit : addBookSubmit)}>
          <div className="login-box">
            <div className="flex-col-gap">
              <InputText
                id="title"
                placeholder="Tytuł"
                {...register('title', {
                  required: true
                })}
                className={errors.title && 'p-invalid'}
              />

              {errors.title?.type === 'required' && (
                <small id="title-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <InputTextarea
                rows={5}
                cols={30}
                autoResize
                id="description"
                {...register('description', {
                  required: true
                })}
                className={errors.description && 'p-invalid'}
                placeholder="Opis"
              />

              {errors.description?.type === 'required' && (
                <small id="description-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <InputText
                id="publisher"
                placeholder="Wydawnictwo"
                {...register('publisher', {
                  required: true
                })}
                className={errors.publisher && 'p-invalid'}
              />

              {errors.publisher?.type === 'required' && (
                <small id="publisher-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <Calendar
                id="dateRelease"
                value={date}
                placeholder="Data wydania"
                {...register('dateRelease', {
                  required: true
                })}
                className={errors.dateRelease && 'p-invalid'}
              />

              {errors.dateRelease?.type === 'required' && (
                <small id="dateRelease-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <Chips
                id="genres"
                placeholder="Gatunki"
                {...register('genres', {
                  required: true
                })}
                value={values}
                onChange={(e) => setValues(e.value)}
                className={errors.genres && 'p-invalid'}
              />

              {errors.genres?.type === 'required' && (
                <small id="genres-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}

              <Chips
                id="authors"
                placeholder="Autorzy"
                {...register('authors', {
                  required: true
                })}
                value={valuesAuthors}
                onChange={(e) => setAuthorsValues(e.value)}
                className={errors.authors && 'p-invalid'}
              />

              {errors.authors?.type === 'required' && (
                <small id="authors-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
              <Controller
                name="quantity"
                control={control}
                rules={{ required: true, min: 1, max: 200000 }}
                render={({ field, fieldState }) => (
                  <InputNumber
                    name={field.name}
                    inputRef={field.ref}
                    placeholder="Ilość książek"
                    buttonLayout="horizontal"
                    step={1}
                    showButtons
                    value={valueAmount}
                    onValueChange={(e) => setValueAmount(e.value)}
                    decrementButtonClassName="p-button-text"
                    incrementButtonClassName="p-button-text"
                    incrementButtonIcon="pi pi-plus"
                    decrementButtonIcon="pi pi-minus"
                    inputClassName={classNames({ 'p-invalid': fieldState.error })}
                  />
                )}
              />

              {errors['quantity'] && (
                <small id="quantity-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
            </div>
          </div>
        </form>
      </Dialog>

      <Button label="Dodaj książkę" onClick={() => setAddBook(true)} className="p-button-text" />
      <div className="card">
        <DataTable value={books} responsiveLayout="scroll">
          <Column field="title" header="Tytuł"></Column>
          <Column field="description" header="Opis"></Column>
          <Column field="quantity" header="Ilość"></Column>
          <Column field="genres" header="Gatunki" body={(data) => joinArray(data.genres)}></Column>
          <Column field="authors" header="Autorzy" body={(data) => joinArray(data.authors)}></Column>
          <Column field="publisher" header="Wydawnictwo"></Column>
          <Column header="Akcje" body={(data) => buttonTemplate(data)}></Column>
        </DataTable>
      </div>
    </div>
  );
}
