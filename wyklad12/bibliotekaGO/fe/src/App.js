import 'primeicons/primeicons.css';
import 'primereact/resources/primereact.min.css';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import { BookSearch } from './components/Book/BookSearch/BookSearch';
import { Login } from './components/User/Login/Login';
import { MainPage } from './components/MainPage/MainPage';
import { Panel } from './components/User/Panel/Panel';
import { Register } from './components/User/Register/Register';
import { Routes, Route } from 'react-router-dom';
import { Error404 } from './components/layout/404/404';
import { BookDetails } from './components/Book/BookDetails/BookDetails';
import { AdminPanel } from './components/Admin/AdminPanel/AdminPanel';
import GuardedRoute from './GuardedRoute';
import React, { useState } from 'react';
import Cookies from 'universal-cookie';
import { useDispatch, useSelector } from 'react-redux';
import { authActions } from './reducers/authReducer';
import { AuthInterceptor } from './AuthInterceptor';

function App() {
  const authenticated = useSelector((state) => state.auth);

  console.log(authenticated);

  return (
    <div>
      <Routes>
        <Route exact path="/" element={<AuthInterceptor />}>
          <Route path="/" element={<GuardedRoute authStatus={authenticated} component={MainPage} />} />
          <Route path="/search" element={<GuardedRoute authStatus={authenticated} component={BookSearch} />} />
          <Route path="/panel" element={<GuardedRoute authStatus={authenticated} component={Panel} />} />
          <Route path="/admin" element={<GuardedRoute authStatus={authenticated} component={AdminPanel} />} />
          <Route path="/details/:id" element={<GuardedRoute authStatus={authenticated} component={BookDetails} />} />
        </Route>

        <Route
          path="/login"
          element={<GuardedRoute authStatus={!authenticated} routeRedirect="/" component={Login} />}
        />
        <Route
          path="/register"
          element={<GuardedRoute authStatus={!authenticated} routeRedirect="/" component={Register} />}
        />

        <Route path="*" element={<Error404 />} />
      </Routes>
    </div>
  );
}

export default App;
