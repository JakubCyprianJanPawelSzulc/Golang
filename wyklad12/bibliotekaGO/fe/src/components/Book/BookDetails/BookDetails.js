import './BookDetails.scss';
import React, { useEffect, useRef, useState } from 'react';
import { NavBar } from '../../layout/NavBar/NavBar';
import { useParams } from 'react-router-dom';
import { booksService } from '../../../services/books.service';
import spring from '../../../assets/images/spring.jpg';
import { Rating } from 'primereact/rating';
import { Toast } from 'primereact/toast';
import { InputTextarea } from 'primereact/inputtextarea';
import { Button } from 'primereact/button';
import { userService } from '../../../services/user.service';
import { Dialog } from 'primereact/dialog';
import { Calendar } from 'primereact/calendar';
import { reservationsService } from '../../../services/reservations.service';

export function BookDetails() {
  let { id } = useParams();
  const [book, setBook] = useState(null);
  const [bookRating, setBookRating] = useState(null);
  const [timeoutTimer, setTimeoutTimer] = useState(null);
  const [bookRated, setBookRated] = useState(null);
  const [comments, setComments] = useState(null);
  const [displayReservationDialog, setDisplayReservationDialog] = useState(false);
  const [reservationDate, setReservationDate] = useState(null);
  const [reservationCount, setReservationCount] = useState(null);

  const commentRef = useRef(null);

  const toast = useRef(null);

  const getData = () => {
    booksService
      .getBookById(id)
      .then((e) => {
        return e;
      })
      .then((ev) => {
        booksService.getBookRateById(id).then((rateNum) => {
          setBook({ ...ev, rate: rateNum.rate, rateCount: rateNum.votes });
        });
      });

    booksService.getUserRatedBook(id).then((e) => {
      const res = e ? e : false;
      setBookRated(res);
    });

    reservationsService.getReservationCountByBookId(id).then((e) => {
      console.log(e);
      setReservationCount(e);
    });
  };

  const getComments = () => {
    booksService.getCommentsByBookId(id).then((e) => {
      setComments(e);
    });
  };

  useEffect(() => {
    getData();
    getComments();
  }, []);

  const sendRate = (rate) => {
    booksService.addRate({ bookId: id, rate: rate }).then((e) => {
      getData();
      toast.current.show({
        severity: 'success',
        summary: 'Zapisano',
        detail: `Ocena została zapisana`,
        life: 3000
      });
    });
  };

  const sendComment = () => {
    const comment = commentRef.current.value;

    if (comment.replace(/\s/g, '').length !== 0 && comment !== null) {
      commentRef.current.value = null;
      booksService.addComment({ bookId: id, comment: comment }).then((e) => {
        getComments();
        toast.current.show({
          severity: 'success',
          summary: 'Wysłano',
          detail: `Komentarz został dodany`,
          life: 3000
        });
      });
    } else {
      toast.current.show({
        severity: 'error',
        summary: 'Zabronione',
        detail: `Puste komentarze są zabronione`,
        life: 3000
      });
    }
  };

  const borrowBook = () => {
    reservationsService.createReservation({ bookId: id, borrowTime: reservationDate.getTime() }).then((e) => {
      console.log(e);
      setDisplayReservationDialog(false);
      getData();
      toast.current.show({
        severity: 'success',
        summary: 'Wypożyczono',
        detail: 'Pomyślnie wypożyczono książkę',
        life: 3000
      });
    });
  };

  const rateBook = (rate) => {
    setBookRating(rate);
    clearTimeout(timeoutTimer);

    toast.current.clear();
    toast.current.show({
      severity: 'info',
      summary: 'Informacja',
      detail: 'Ocena zostanie zapisana za 5 sekund',
      life: 5000
    });

    setTimeoutTimer(
      setTimeout(() => {
        sendRate(rate);
      }, 5000)
    );
  };

  const reservationDialogFooter = (
    <div>
      <Button label="Anuluj" className="p-button-text" onClick={() => setDisplayReservationDialog(false)} />
      <Button className="p-button-rounded" onClick={() => borrowBook()} label="Wypożycz" />
    </div>
  );

  let today = new Date();
  let month = today.getMonth();
  let year = today.getFullYear();

  let nextMonth = month === 11 ? 0 : month + 1;
  let nextYear = nextMonth === 0 ? year + 1 : year;

  let minDate = new Date();

  let maxDate = new Date();
  maxDate.setMonth(nextMonth);
  maxDate.setFullYear(nextYear);

  return (
    <div>
      <Dialog
        header={`Wypożyczenie książki: ${book?.title}`}
        closable={false}
        visible={displayReservationDialog}
        style={{ width: '50vw' }}
        footer={reservationDialogFooter}
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

      <Toast ref={toast} />
      <NavBar></NavBar>
      {book && (
        <div className="details">
          <div className="photo">
            <img src={spring} />
          </div>
          <div className="right">
            <div className="label">Tytuł</div>
            <div className="value">{book.title}</div>
            <div className="label">Opis</div>
            <div className="value">{book.description}</div>
            <div className="label">Autorzy</div>
            <div className="value">{book.authors.join(', ')}</div>
            <div className="label">Data wydania</div>
            <div className="value">
              {new Date(parseInt(book.dateRelease)).toLocaleString('pl-PL', {
                year: 'numeric',
                month: 'numeric',
                day: 'numeric'
              })}
            </div>
            <div className="label">Gatunki</div>
            <div className="value">{book.genres.join(', ')}</div>
            <div className="label">Wydawnictwo</div>
            <div className="value">{book.publisher}</div>
            <div className="label">Średnia ocena</div>
            <div className="value">
              {book?.rate || '-'} {book?.rateCount ? '( ' + book?.rateCount + ' )' : ''}
            </div>

            <div className="label">Twoja ocena</div>
            <div className="value">
              {bookRated !== null && (
                <Rating
                  value={bookRated ? bookRated.rate : bookRating}
                  readOnly={bookRated}
                  cancel={false}
                  onChange={(e) => rateBook(e.value)}
                />
              )}
            </div>

            <div className={reservationCount >= book.quantity ? 'span-2' : ''}>
              {reservationCount !== null && reservationCount < book.quantity && (
                <Button
                  label="Wypożycz"
                  onClick={() => setDisplayReservationDialog(true)}
                  className=" p-button-outlined p-button-help"
                />
              )}
              {reservationCount >= book.quantity && <div> Wszystkie egzemplarze są aktualnie wypożyczone.</div>}
            </div>
          </div>
        </div>
      )}

      <div className="outer">
        <div className="comment">
          <div className="top-c">
            <InputTextarea rows={5} ref={commentRef} cols={30} autoResize />
            <Button label="Wyślij" onClick={sendComment} className="p-button-text" />
          </div>

          <div className="comments">
            {comments &&
              comments.map((e) => {
                return (
                  <div key={e._id} className="comm">
                    <div className="name">
                      {e.userName}&nbsp;{e.userSurName}
                    </div>
                    <div className="cont-com">{e.comment}</div>
                  </div>
                );
              })}
            {comments && comments.length === 0 && <div>Brak komentarzy.</div>}
          </div>
        </div>
      </div>
    </div>
  );
}
