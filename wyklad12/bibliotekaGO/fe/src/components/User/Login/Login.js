import './Login.scss';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { useForm } from 'react-hook-form';
import sha512 from 'crypto-js/sha512';
import { useNavigate } from 'react-router-dom';
import { userService } from '../../../services/user.service';
import React, { useRef, useState } from 'react';
import Cookies from 'universal-cookie';
import { authActions } from '../../../reducers/authReducer';
import { useDispatch, useSelector } from 'react-redux';

export function Login() {
  const ref = useRef(null);
  const navigate = useNavigate();
  const [loginSuccess, setLoginSuccess] = useState(false);
  const cookies = new Cookies();

  const dispatch = useDispatch();

  const {
    control,
    register,
    handleSubmit,
    setError,
    formState: { errors }
  } = useForm();

  const loginUser = (data) => {
    const hashedPass = sha512(data.pass).toString();
    data = { ...data, pass: hashedPass };

    userService.loginUser(data).then((e) => {
      console.log(e);
      if (e.status === 200) {
        const date = new Date();
        const tomorrow = date.setDate(date.getDate() + 1);
        console.log(tomorrow);
        cookies.set('loginCookie', e.data, { path: '/', expires: new Date(tomorrow) });
        console.log(cookies.getAll());
        setLoginSuccess(true);
        ref.current.classList.remove('p-button-danger');
        ref.current.classList.add('p-button-success');

        setTimeout(() => {
          dispatch(authActions.authenticate());
          navigate('/');
        }, 2000);
      } else {
        ref.current.classList.remove('p-button-success');
        ref.current.classList.add('p-button-danger');
      }
    });
  };

  return (
    <div className="login">
      <h2>Login</h2>
      <form onSubmit={handleSubmit(loginUser)}>
        <div className="login-box">
          <div className="flex-col-gap">
            <InputText
              id="email"
              placeholder="Email"
              {...register('email', {
                required: true,
                pattern: {
                  value:
                    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
                }
              })}
              className={errors.email && 'p-invalid'}
              disabled={loginSuccess}
            />

            {errors.email?.type === 'required' && (
              <small id="email-req" className="p-error errMessage">
                Pole jest wymagane.
              </small>
            )}

            {errors.email?.type === 'pattern' && (
              <small id="email-req" className="p-error errMessage">
                Proszę wpisać poprawny email.
              </small>
            )}
          </div>
          <div className="flex-col-gap">
            <InputText
              {...register('pass', { required: true })}
              className={errors.pass && 'p-invalid'}
              id="pass"
              type="password"
              placeholder="Hasło"
              disabled={loginSuccess}
            />

            {errors.pass?.type === 'required' && (
              <small id="pass-req" className="p-error errMessage">
                Pole jest wymagane.
              </small>
            )}
          </div>
          <Button ref={ref} label="Zaloguj" className="p-button-rounded p-button-info" />
          <div className="login-redirect" onClick={() => navigate('/register')}>
            Rejestracja
          </div>
        </div>
      </form>
    </div>
  );
}
