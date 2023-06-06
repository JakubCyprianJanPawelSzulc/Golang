import React from 'react';
import { Route, Navigate } from 'react-router-dom';

const GuardedRoute = ({ authStatus, component: Component, routeRedirect }) => {
  if (!authStatus) {
    if (!routeRedirect) {
      return <Navigate to="/login" replace />;
    } else {
      return <Navigate to={routeRedirect} replace />;
    }
  }

  return <Component />;
};

export default GuardedRoute;
