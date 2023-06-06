import { NavBar } from '../../layout/NavBar/NavBar';
import './AdminPanel.scss';
import { BookActions } from '../BookActions/BookActions';
import React, { useState, useEffect, useRef } from 'react';
import { StatisticsActions } from '../StatisticsActions/StatisticsActions';
import { CommentActions } from '../CommentActions/CommentActions';
import { UserActions } from '../UserActions/UserActions';
import { ReservationActions } from '../ReservationActions/ReservationActions';

export function AdminPanel() {
  const [activeSection, setActiveSection] = useState(1);

  return (
    <div className="admin-panel">
      <NavBar></NavBar>

      <div className="nav-section">
        <div className={['link', activeSection === 1 && 'active'].join(' ')} onClick={() => setActiveSection(1)}>
          Książki
        </div>
        <div className={['link', activeSection === 2 && 'active'].join(' ')} onClick={() => setActiveSection(2)}>
          Użytkownicy
        </div>
        <div className={['link', activeSection === 3 && 'active'].join(' ')} onClick={() => setActiveSection(3)}>
          Komentarze
        </div>

        <div className={['link', activeSection === 4 && 'active'].join(' ')} onClick={() => setActiveSection(4)}>
          Rezerwacje
        </div>
        <div className={['link', activeSection === 5 && 'active'].join(' ')} onClick={() => setActiveSection(5)}>
          Statystyki
        </div>
      </div>

      <div className="dynamic-content">
        {activeSection === 1 && <BookActions></BookActions>}
        {activeSection === 2 && <UserActions></UserActions>}
        {activeSection === 3 && <CommentActions></CommentActions>}
        {activeSection === 4 && <ReservationActions></ReservationActions>}
        {activeSection === 5 && <StatisticsActions></StatisticsActions>}
      </div>
    </div>
  );
}
