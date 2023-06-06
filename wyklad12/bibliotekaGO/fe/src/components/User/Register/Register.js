import './Register.scss';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { InputMask } from 'primereact/inputmask';
import React, { useRef } from 'react';
import { useForm, Controller } from 'react-hook-form';
import { Password } from 'primereact/password';
import { Divider } from 'primereact/divider';
import { userService } from '../../../services/user.service';
import sha512 from 'crypto-js/sha512';
import { useNavigate } from 'react-router-dom';
import { Toast } from 'primereact/toast';

export function Register() {
  const toast = useRef(null);
  const navigate = useNavigate();

  const {
    control,
    register,
    reset,
    handleSubmit,
    setError,
    formState: { errors }
  } = useForm();

  const registerUser = (data) => {
    const hashedPass = sha512(data.pass).toString();
    data = { ...data, pass: hashedPass };

    userService.registerUser(data).then((e) => {
      reset();
      toast.current.show({
        severity: 'success',
        summary: 'Sukces',
        detail: 'Pomyślnie zarejestrowano.',
        life: 5000
      });
    });
  };

  const header = <h6 className="header-pas">Wybierz hasło</h6>;
  const footer = (
    <div>
      <Divider />
      <div className="password-footer">
        <div className="sug">Sugestie</div>

        <ul style={{ lineHeight: '1.5' }}>
          <li>Przynajmniej jedna mała litera</li>
          <li>Przynajmniej jedna duża litera</li>
          <li>Przynajmniej jedna cyfra</li>
          <li>Minimum 8 znaków</li>
        </ul>
      </div>
    </div>
  );

  return (
    <div className="register">
      <Toast ref={toast} />
      <h2>Rejestracja</h2>
      <form onSubmit={handleSubmit(registerUser)}>
        <div className="register-box">
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
            <Controller
              name="pass"
              control={control}
              defaultValue={''}
              rules={{
                required: 'Pole jest wymagane.'
              }}
              render={({ field, fieldState }) => (
                <Password
                  {...field}
                  inputRef={field.ref}
                  type="password"
                  placeholder="Hasło"
                  header={header}
                  weakLabel="Słabe"
                  mediumLabel="Średnie"
                  strongLabel="Silne"
                  footer={footer}
                  promptLabel=" "
                  className={errors['pass'] && 'p-invalid'}
                />
              )}
            />
            {errors['pass'] && (
              <small id="password-req" className="p-error errMessage">
                Pole jest wymagane.
              </small>
            )}
          </div>

          <div className="flex-col-gap">
            <InputText
              id="name"
              placeholder="Imię"
              {...register('name', { required: true })}
              className={errors.name && 'p-invalid'}
            />

            {errors.name?.type === 'required' && (
              <small id="name-req" className="p-error errMessage">
                Pole jest wymagane.
              </small>
            )}
          </div>

          <div className="flex-col-gap">
            <InputText
              id="surName"
              placeholder="Nazwisko"
              {...register('surName', { required: true })}
              className={errors.surName && 'p-invalid'}
            />

            {errors.surName?.type === 'required' && (
              <small id="surName-req" className="p-error errMessage">
                Pole jest wymagane.
              </small>
            )}
          </div>

          <div className="half">
            <div className="flex-col-gap inline">
              <InputMask
                mask="99-999"
                placeholder="00-000"
                id="postal"
                {...register('postal', { required: true })}
                className={errors.postal && 'p-invalid'}
              ></InputMask>

              {errors.postal?.type === 'required' && (
                <small id="postal-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
            </div>
            <div className="flex-col-gap inline">
              <InputText
                id="city"
                placeholder="Miasto"
                {...register('city', { required: true })}
                className={errors.city && 'p-invalid'}
              />

              {errors.city?.type === 'required' && (
                <small id="city-req" className="p-error errMessage">
                  Pole jest wymagane.
                </small>
              )}
            </div>
          </div>

          <Button label="Zarejestruj" className="p-button-rounded p-button-info" />

          <div className="login-redirect" onClick={() => navigate('/login')}>
            Login
          </div>
        </div>
      </form>
    </div>
  );
}
