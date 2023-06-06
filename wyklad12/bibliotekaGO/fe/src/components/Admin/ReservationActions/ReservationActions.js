import './ReservationActions.scss';
import React, { useState, useEffect, useRef } from 'react';
import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';
import { reservationsService } from '../../../services/reservations.service';

export function ReservationActions() {
  const [reservations, setReservations] = useState(null);
  const toast = useRef(null);

  useEffect(() => {
    getReservations();
  }, []);

  const getReservations = () => {
    reservationsService.getAllReservations().then((e) => {
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

  const newActions = (rowData) => {
    return (
      <div className="admin-actions">
        {rowData.status !== 'CANCELLED' && rowData.status !== 'RETURNED' && (
          <Button
            label="Następny status"
            onClick={() => nextStep(rowData)}
            className="p-button-rounded "
            aria-label="Następny status"
          />
        )}
      </div>
    );
  };

  const nextStep = (data) => {
    const time = new Date().getTime();
    switch (data.status) {
      //wypozyczenie
      case 'READY':
        reservationsService.editGivenReservation({ reservationId: data._id, handOverTime: time }).then((e) => {
          getReservations();
        });

        break;

      //zwrot
      case 'BORROWED':
        reservationsService.editReturnedReservation({ reservationId: data._id, returnTime: time }).then((e) => {
          getReservations();
        });
        break;

      default:
        break;
    }

    reservationsService.updateReservationStatus(data._id).then((e) => {
      getReservations();

      toast.current.show({
        severity: 'success',
        summary: 'Zaktualizowano',
        detail: `Pomyślnie zaktualizowano status`,
        life: 3000
      });
    });
  };

  const getFullName = (rowData) => {
    return rowData.userName + ' ' + rowData.userSurName;
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

  return (
    <div className="page">
      <Toast ref={toast} />

      <div className="content">
        <div className="card admin">
          <DataTable value={reservations} header="Wypożyczenia" responsiveLayout="scroll">
            <Column field="bookTitle" header="Tytuł" />
            <Column body={getFullName} header="Użytkownik" />
            <Column field="borrowTime" body={(e) => transformToDate(e, 'borrowTime')} header="Zarezerwowano do" />
            <Column field="handOverTime" body={(e) => transformToDate(e, 'handOverTime')} header="Wydano" />
            <Column field="returnTime" body={(e) => transformToDate(e, 'returnTime')} header="Zwrócono" />
            <Column field="name" header="Łączny czas wypożyczenia" body={calculateResTime} />
            <Column field="status" header="Status" body={statusTemplate} />
            <Column header="Akcje" body={newActions} />
          </DataTable>
        </div>
      </div>
    </div>
  );
}
