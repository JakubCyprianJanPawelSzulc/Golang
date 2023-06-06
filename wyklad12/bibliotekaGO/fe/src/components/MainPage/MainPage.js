import './MainPage.scss';
import { NavBar } from '../layout/NavBar/NavBar';
import author from '../../assets/images/author.jpg';
import { useNavigate } from 'react-router-dom';
import { booksService } from '../../services/books.service';
import { useEffect, useState } from 'react';

export function MainPage() {
  const navigate = useNavigate();

  const [books, setBooks] = useState([]);

  useEffect(() => {
    Promise.all([booksService.getAllBooks(), booksService.getAllRatesCountByBookId()]).then((values) => {
      console.log(values);
      setBooks(
        values[0].data
          .map((e) => {
            return { ...e, ...values[1].data?.find((el) => el._id === e._id) };
          })
          .sort((a, b) => {
            const bRate = b.rate ? b.rate : 0;
            const aRate = a.rate ? a.rate : 0;
            return parseInt(bRate) - parseInt(aRate);
          })
      );
    });
  }, []);

  useEffect(() => {
    console.log(books);
  }, [books]);
  return (
    <div>
      <NavBar></NavBar>
      <div className="content">
        <div className="main">
          {books &&
            books.map((e) => {
              return (
                <div key={e._id} className="block" onClick={() => navigate('/details/' + e._id)}>
                  <div className="cover">
                    <div className="photo">
                      <img src={author} alt="author" />
                    </div>
                    <div className="right">
                      <div className="title">{e.title}</div>
                      <div className="description">{e.description}</div>
                    </div>
                  </div>
                </div>
              );
            })}
        </div>
      </div>
    </div>
  );
}
