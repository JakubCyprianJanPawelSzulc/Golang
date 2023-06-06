import axios from 'axios';
import { backendURL } from '../constants/backend.constant';

async function createReservation(reservationData) {
  return axios
    .post(backendURL + '/reservations/createReservation', reservationData)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function editReservation(reservationData) {
  return axios
    .post(backendURL + '/reservations/editReservation', reservationData)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getReservationCountByBookId(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .get(backendURL + '/reservations/getReservationCountByBookId', { params: params })
    .then(function (response) {
      return response.data.count;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getUserReservations() {
  return axios
    .get(backendURL + '/reservations/getUserReservations')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getAllReservations() {
  return axios
    .get(backendURL + '/reservations/getAllReservations')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function payReservation(reservationId) {
  const params = new URLSearchParams([['reservationId', reservationId]]);
  return axios
    .get(backendURL + '/reservations/payReservation', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function updateReservationStatus(reservationId) {
  const params = new URLSearchParams([['reservationId', reservationId]]);
  return axios
    .get(backendURL + '/reservations/updateReservationStatus', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function cancelReservation(reservationId) {
  const params = new URLSearchParams([['reservationId', reservationId]]);
  return axios
    .get(backendURL + '/reservations/cancelReservation', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function editGivenReservation(reservationData) {
  return axios
    .post(backendURL + '/reservations/editGivenReservation', reservationData)
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function editReturnedReservation(reservationData) {
  return axios
    .post(backendURL + '/reservations/editReturnedReservation', reservationData)
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

export const reservationsService = {
  createReservation,
  getReservationCountByBookId,
  payReservation,
  updateReservationStatus,
  getUserReservations,
  editReservation,
  cancelReservation,
  getAllReservations,
  editGivenReservation,
  editReturnedReservation
};
