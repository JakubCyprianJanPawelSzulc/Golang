import './BookSearch.scss';
import React, { useEffect, useState } from 'react';
import { NavBar } from '../../layout/NavBar/NavBar';
import { InputText } from 'primereact/inputtext';
import author from '../../../assets/images/author.jpg';
import { useNavigate } from 'react-router-dom';
import { booksService } from '../../../services/books.service';
import { Dropdown } from 'primereact/dropdown';

export function BookSearch() {
  const navigate = useNavigate();

  const [books, setBooks] = useState([]);
  const [searchPhrase, setSearchPhrase] = useState(null);
  const [sort, setSort] = useState(null);

  useEffect(() => {
    if (searchPhrase) {
      booksService.getAllBooks().then((e) => {
        if (sort === null) {
          setBooks(e.data);
          console.log(e.data);
        } else if (sort) {
          setBooks(
            e.data.sort((a, b) => {
              return a.dateRelease - b.dateRelease;
            })
          );
        } else {
          setBooks(
            e.data.sort((a, b) => {
              return b.dateRelease - a.dateRelease;
            })
          );
        }
      });
    }
  }, [searchPhrase, sort]);

  const sortOptions = [
    { name: 'Najnowsze', value: true },
    { name: 'Najstarsze', value: false }
  ];
  return (
    <div className="page">
      <NavBar></NavBar>
      <div className="search-bar">
        <span className="p-float-label">
          <InputText id="username" onChange={(e) => setSearchPhrase(e.target.value)} />
          <label htmlFor="username">Szukaj książki</label>
        </span>
        <Dropdown
          value={sort}
          onChange={(e) => setSort(e.target.value)}
          options={sortOptions}
          optionLabel="name"
          placeholder="Sortuj według"
        />
      </div>

      <div className="content">
        <div className="main">
          {books &&
            books
              .filter(
                (item) =>
                  item.title.toLowerCase().includes(searchPhrase.toLowerCase()) ||
                  item.authors.some((e) => e.toLowerCase().includes(searchPhrase.toLowerCase())) ||
                  item.genres.some((e) => e.toLowerCase().includes(searchPhrase.toLowerCase()))
              )
              .map((e) => {
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
