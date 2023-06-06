import { NavBar } from '../../layout/NavBar/NavBar';
import './Panel.scss';
import React, { useState, useEffect, useRef } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Rating } from 'primereact/rating';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { Dialog } from 'primereact/dialog';
import { Calendar } from 'primereact/calendar';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';
import { InputMask } from 'primereact/inputmask';
import { useForm, Controller } from 'react-hook-form';
import { ProgressSpinner } from 'primereact/progressspinner';
import { reservationsService } from '../../../services/reservations.service';

export function Panel() {
  const [reservations, setReservations] = useState(null);
  const [payReservationVisible, setPayReservationVisible] = useState(false);
  const [editReservationVisible, setEditReservationVisible] = useState(false);
  const [payProceed, setPayProceed] = useState(false);
  const [reservationDate, setReservationDate] = useState(null);
  const [blikValue, setBlikValue] = useState(null);
  const toast = useRef(null);
  const [contextRow, setContextRow] = useState(null);

  let today = new Date();
  let month = today.getMonth();
  let year = today.getFullYear();

  let nextMonth = month === 11 ? 0 : month + 1;
  let nextYear = nextMonth === 0 ? year + 1 : year;

  let minDate = new Date();

  let maxDate = new Date();
  maxDate.setMonth(nextMonth);
  maxDate.setFullYear(nextYear);

  useEffect(() => {
    getReservations();
  }, []);

  const getReservations = () => {
    reservationsService.getUserReservations().then((e) => {
      console.log(e);
      setReservations(e);
    });
  };

  const transformToDate = (rowData, propName) => {
    if (rowData[propName]) {
      return new Date(parseInt(rowData[propName])).toLocaleString('pl-PL', {
        year: 'numeric',
        month: 'numeric',
        day: 'numeric'
      });
    } else {
      return '-';
    }
  };

  const editBorrow = () => {
    reservationsService
      .editReservation({ reservationId: contextRow._id, borrowTime: reservationDate.getTime() })
      .then((e) => {
        setEditReservationVisible(false);
        getReservations();
        toast.current.show({
          severity: 'success',
          summary: 'Edytowano',
          detail: 'Pomyślnie edytowano rezerwację',
          life: 3000
        });
      });
  };

  const reservationEditDialogFooter = (
    <div>
      <Button
        label="Anuluj"
        className="p-button-rounded p-button-text"
        onClick={() => setEditReservationVisible(false)}
      />
      <Button className="p-button-rounded" onClick={() => editBorrow()} label="Edytuj" />
    </div>
  );

  const statusTemplate = (rowData) => {
    let statusTranslated;

    switch (rowData.status) {
      case 'NEW':
        statusTranslated = 'NOWE';
        break;
      case 'CONFIRMED':
        statusTranslated = 'OPŁACONE';
        break;
      case 'READY':
        statusTranslated = 'GOTOWE';
        break;
      case 'BORROWED':
        statusTranslated = 'WYPOŻYCZONE';
        break;
      case 'RETURNED':
        statusTranslated = 'ODDANE';
        break;
      case 'CANCELLED':
        statusTranslated = 'ANULOWANE';
        break;
      default:
        statusTranslated = '-';
        break;
    }
    return (
      <span
        className={`status-badge status-${rowData.status ? rowData.status.toLowerCase().replace(/\s+/g, '_') : ''}`}
      >
        {statusTranslated}
      </span>
    );
  };

  const ratingTemplate = (rowData) => {
    return <Rating value={rowData.ratedValue} readOnly cancel={false} />;
  };

  const acceptCancel = (data) => {
    reservationsService.cancelReservation(data._id).then((e) => {
      toast.current.show({
        severity: 'success',
        summary: 'Anulowano',
        detail: 'Pomyślnie anulowano rezerwację.',
        life: 3000
      });

      getReservations();
    });
  };

  const confirmCancel = (rowData) => {
    confirmDialog({
      message: 'Jesteś pewnien, że chcesz usunąć tę książkę?',
      header: 'Potwierdzenie usunięcia',
      icon: 'pi pi-exclamation-triangle',
      acceptLabel: 'Tak',
      rejectLabel: 'Nie',
      acceptClassName: 'p-button-danger',
      accept: () => acceptCancel(rowData),
      reject: () => {}
    });
  };

  const calculateResTime = (rowData) => {
    if (rowData.returnTime) {
      const day = 24 * 60 * 60 * 1000;
      const f1 = new Date(rowData.handOverTime);
      const f2 = new Date(rowData.returnTime);

      const diff = Math.round(Math.abs((f1 - f2) / day));

      return diff + ' dni';
    } else {
      return '-';
    }
  };

  const payProceedAction = () => {
    if (blikValue?.length > 5) {
      setPayProceed(true);

      setTimeout(() => {
        const update = reservationsService.updateReservationStatus(contextRow._id);
        const pay = reservationsService.payReservation(contextRow._id);
        Promise.all([update, pay]).then((values) => {
          setPayProceed(false);
          setPayReservationVisible(false);
          setBlikValue(null);
          getReservations();

          toast.current.show({
            severity: 'success',
            summary: 'Opłacono',
            detail: `Sesja została pomyślnie opłacona`,
            life: 5000
          });
        });
      }, 3000);
    }
  };

  const showPayModal = (rowData) => {
    setContextRow(rowData);
    setPayReservationVisible(true);
  };

  const showEditModal = (rowData) => {
    setContextRow(rowData);
    setEditReservationVisible(true);
  };

  const cancelPay = () => {
    setPayReservationVisible(false);
    setBlikValue(null);
  };

  const newActions = (rowData) => {
    if (rowData.status === 'NEW') {
      return (
        <div className="newActions">
          <Button
            icon="pi pi-money-bill"
            onClick={() => showPayModal(rowData)}
            className="p-button-rounded p-button-success"
            aria-label="Opłać"
          />
          <Button
            icon="pi pi-pencil"
            onClick={() => showEditModal(rowData)}
            className="p-button-rounded p-button-secondary"
            aria-label="Edytuj"
          />
          <Button
            icon="pi pi-times"
            onClick={() => confirmCancel(rowData)}
            className="p-button-rounded p-button-danger"
            aria-label="Anuluj"
          />
        </div>
      );
    }
  };
  const payFooter = (
    <div>
      <Button
        label="Anuluj"
        disabled={payProceed}
        className=" p-button-rounded p-button-text"
        onClick={() => cancelPay()}
      />
      <Button label="Opłać" disabled={payProceed} onClick={() => payProceedAction()} className="p-button-rounded" />
    </div>
  );
  return (
    <div className="page">
      <NavBar></NavBar>
      <ConfirmDialog />
      <Toast ref={toast} />

      <Dialog
        header="Opłać rezerwację"
        footer={payFooter}
        visible={payReservationVisible}
        style={{ width: '50vw' }}
        closable={false}
        modal
      >
        <div className="pay-dialog">
          {!payProceed && (
            <div>
              <div className="header">Wpisz kod blik:</div>
              <InputMask
                value={blikValue}
                onChange={(e) => setBlikValue(e.value)}
                mask="999-999"
                placeholder="000-000"
              ></InputMask>
            </div>
          )}
          {payProceed && <ProgressSpinner />}
        </div>
      </Dialog>

      <Dialog
        header={`Edycja rezerwacji: ${contextRow?.bookTitle}`}
        closable={false}
        visible={editReservationVisible}
        style={{ width: '50vw' }}
        footer={reservationEditDialogFooter}
      >
        <div className="dialog-reservation">
          <Calendar
            id="reservationReturnDate"
            placeholder="Do kiedy chcesz wypożyczyć?"
            onChange={(e) => setReservationDate(e.value)}
            minDate={minDate}
            maxDate={maxDate}
            readOnlyInput
          />
        </div>
      </Dialog>

      <div className="top">
        <div className="data-box">
          <div className="icon">
            <i className="pi pi-money-bill"></i>
          </div>
          <div className="value">${reservations && reservations.length * 5}</div>
          <div className="desc">Zaoszczędzonych pieniędzy</div>
        </div>
        <div className="data-box">
          <div className="icon">
            <i className="pi pi-bolt"></i>
          </div>
          <div className="value">{reservations && reservations.length}</div>
          <div className="desc">Wypożyczone książki</div>
        </div>
        <div className="data-box">
          <div className="icon">
            <i className="pi pi-check-circle"></i>
          </div>
          <div className="value">0</div>
          <div className="desc">Zaległości</div>
        </div>
      </div>

      <div className="content">
        <div className="card">
          <DataTable value={reservations} header="Wypożyczenia" responsiveLayout="scroll">
            <Column field="bookTitle" header="Tytuł" />
            <Column field="borrowTime" body={(e) => transformToDate(e, 'borrowTime')} header="Zarezerwowano do" />
            <Column field="handOverTime" body={(e) => transformToDate(e, 'handOverTime')} header="Wydano" />
            <Column field="returnTime" body={(e) => transformToDate(e, 'returnTime')} header="Zwrócono" />
            <Column field="name" header="Łączny czas wypożyczenia" body={calculateResTime} />

            <Column field="status" header="Status" body={statusTemplate} />
            <Column field="rating" header="Ocena" body={ratingTemplate} />
            <Column header="Akcje" body={newActions} />
          </DataTable>
        </div>
      </div>
    </div>
  );
}
