import axios from 'axios';
import { backendURL } from '../constants/backend.constant';

async function getAllBooks() {
  return axios
    .get(backendURL + '/books/getAllBooks')
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getAllRatesByBookId() {
  return axios
    .get(backendURL + '/books/getAllRatesByBookId')
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getAllRatesCountByBookId() {
  return axios
    .get(backendURL + '/books/getAllRatesCountByBookId')
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getFiveMostPopularBooks() {
  return axios
    .get(backendURL + '/books/getFiveMostPopularBooks')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getAllComments() {
  return axios
    .get(backendURL + '/books/getAllComments')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getBookById(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .get(backendURL + '/books/getBookById', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getBookRateById(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .get(backendURL + '/books/getBookRateById', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getCommentsByBookId(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .get(backendURL + '/books/getCommentsByBookId', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function deleteBookbyId(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .delete(backendURL + '/books/deleteBookbyId', { params: params })
    .then(function (response) {
      return response.data.rate;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function addRate(rate) {
  return axios
    .post(backendURL + '/books/addRate', rate)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function addBook(book) {
  return axios
    .post(backendURL + '/books/addBook', book)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function addComment(comment) {
  return axios
    .post(backendURL + '/books/addComment', comment)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getUserRatedBook(bookId) {
  const params = new URLSearchParams([['bookId', bookId]]);
  return axios
    .get(backendURL + '/books/getUserRatedBook', { params: params })
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getTenMostBooks() {
  return axios
    .get(backendURL + '/books/getTenMostBooks')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function getTenOldestBooks() {
  return axios
    .get(backendURL + '/books/getTenOldestBooks')
    .then(function (response) {
      return response.data;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function updateBook(book) {
  return axios
    .put(backendURL + '/books/updateBook', book)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function updateComment(comment) {
  return axios
    .put(backendURL + '/books/updateComment', comment)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function deleteCommentbyId(commentId) {
  const params = new URLSearchParams([['commentId', commentId]]);
  return axios
    .delete(backendURL + '/books/deleteCommentbyId', { params: params })
    .then(function (response) {
      return response.data.rate;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function addCommentByAdmin(comment) {
  return axios
    .post(backendURL + '/books/addCommentByAdmin', comment)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

async function addPhoto(data) {
  return axios
    .post(backendURL + '/books/addPhoto', data)
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
      return error;
    });
}

export const booksService = {
  getAllBooks,
  getBookById,
  getBookRateById,
  addBook,
  updateBook,
  deleteBookbyId,
  getAllRatesByBookId,
  getUserRatedBook,
  addRate,
  getAllRatesCountByBookId,
  addComment,
  getCommentsByBookId,
  getTenMostBooks,
  getTenOldestBooks,
  getFiveMostPopularBooks,
  getAllComments,
  deleteCommentbyId,
  updateComment,
  addCommentByAdmin,
  addPhoto
};
