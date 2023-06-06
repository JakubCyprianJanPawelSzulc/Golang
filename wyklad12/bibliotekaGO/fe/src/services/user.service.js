import axios from 'axios';
import { backendURL } from '../constants/backend.constant';

async function registerUser(userData) {
  axios
    .post(backendURL + '/register', userData)
    .then(function (response) {
      console.log('RESPONSE: ', response.data);
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function loginUser(userData) {
  return axios
    .post(backendURL + '/login', userData)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function logoutUser(userData) {
  return axios
    .post(backendURL + '/login/logout', userData)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getAllUsers() {
  return axios
    .get(backendURL + '/user/getAllUsers')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function deleteUserbyId(userId) {
  const params = new URLSearchParams([['userId', userId]]);
  return axios
    .get(backendURL + '/user/deleteUserbyId', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function editUserNameAndSurname(userNameAndSurname) {
  return axios
    .post(backendURL + '/user/editUserNameAndSurname', userNameAndSurname)
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

export const userService = { registerUser, loginUser, logoutUser, getAllUsers, deleteUserbyId, editUserNameAndSurname };
