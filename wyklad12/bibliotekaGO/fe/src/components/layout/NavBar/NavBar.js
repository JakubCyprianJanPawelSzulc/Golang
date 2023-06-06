import './NavBar.scss';
import React, { useLayoutEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import { authActions } from '../../../reducers/authReducer';
import { useDispatch, useSelector } from 'react-redux';

export function NavBar() {
  const dispatch = useDispatch();
  const SlideRef = React.useRef(null);
  const LogOutRef = React.useRef(null);
  const navigate = useNavigate();
  const location = useLocation();

  const [expanded, setExpanded] = useState(false);

  const changeExpanded = () => {
    setExpanded(!expanded);
  };

  useLayoutEffect(() => {
    changeslide();
  }, []);

  const logout = () => {
    dispatch(authActions.DeAuthenticate());
  };

  const hide = () => {
    LogOutRef.current.classList.remove('slide-in-logout');
    LogOutRef.current.classList.add('slide-out-logout');
  };

  const show = () => {
    LogOutRef.current.classList.remove('invisible');
    LogOutRef.current.classList.add('slide-in-logout');
    LogOutRef.current.classList.remove('slide-out-logout');
  };

  const changeslide = () => {
    SlideRef.current.style.height = expanded ? '60px' : 0;
    expanded ? show() : hide();
    changeExpanded();
  };

  return (
    <div>
      <div className="navbar">
        <div className="logo" onClick={changeslide}>
          <i className="pi pi-align-center"></i>
        </div>

        <div className="text">Books</div>

        <div className="logo smaller">
          <i className="pi pi-sign-out absolute invisible" ref={LogOutRef} onClick={logout}></i>
        </div>
      </div>
      <div className="slide-in" ref={SlideRef}>
        {location.pathname !== '/' && (
          <div className="page" onClick={() => navigate('/')}>
            Główna
          </div>
        )}

        {location.pathname !== '/search' && (
          <div className="page" onClick={() => navigate('/search')}>
            Szukaj
          </div>
        )}

        {location.pathname !== '/panel' && (
          <div className="page" onClick={() => navigate('/panel')}>
            Panel
          </div>
        )}
      </div>
    </div>
  );
}
