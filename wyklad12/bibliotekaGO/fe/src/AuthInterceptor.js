import axios from 'axios';
import { useDispatch } from 'react-redux';
import { Outlet } from 'react-router-dom';
import Cookies from 'universal-cookie';
import { authActions } from './reducers/authReducer';

const AuthInterceptor = () => {
  const cookies = new Cookies();
  const dispatch = useDispatch();

  axios.interceptors.request.use(
    function (config) {
      config.headers.userToken = cookies.get('loginCookie')?.token;
      return config;
    },
    function (error) {
      return Promise.reject(error);
    }
  );

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response?.status === 401) {
        //Api returns Unauthorized , logout
        dispatch(authActions.DeAuthenticate());
      }
    }
  );

  return <Outlet />;
};
export { AuthInterceptor };
