import { createSlice } from '@reduxjs/toolkit';
import Cookies from 'universal-cookie';
const cookies = new Cookies();
import { userService } from '../services/user.service';

const authSlice = createSlice({
  name: 'auth',
  initialState: cookies.get('loginCookie') ? true : false,
  reducers: {
    authenticate: (state, action) => {
      return true;
    },
    DeAuthenticate: (state, action) => {
      userService.logoutUser({ userToken: cookies.get('loginCookie')?.token });
      cookies.remove('loginCookie');
      return false;
    }
  }
});

export const authActions = authSlice.actions;
export default authSlice.reducer;
