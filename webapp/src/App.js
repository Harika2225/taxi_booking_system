import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { NavBar } from './components/NavBar';
import { Banner } from './components/Banner';
import DocsPage from './components/Docs';
import SwaggerUI from 'swagger-ui-react';
import 'swagger-ui-react/swagger-ui.css';
import PrivateRoute from './config/auth/privateRoute';
import NotesList from './components/notes/NotesList';

function App() {
  return (
    <div className="App">
      <NavBar />
      <Router>
        <Routes>
          <Route path="/" element={<Banner />} />
          <Route
            path="/docs"
            element={
              <PrivateRoute>
                <div className="container">
                  <DocsPage />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/bookingmanagementmodule"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_BOOKINGMANAGEMENTMODULE.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/customermanagementmodule"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_CUSTOMERMANAGEMENTMODULE.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/drivermanagementmodule"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_DRIVERMANAGEMENTMODULE.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/swagger/paymentmanagementmodule"
            element={
              <PrivateRoute>
                <div className="swagger">
                  <SwaggerUI url={process.env.REACT_APP_MICROSERVICE_PAYMENTMANAGEMENTMODULE.concat('/v3/api-docs')} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/bookingmanagementmodule"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'bookingmanagementmodule'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/customermanagementmodule"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'customermanagementmodule'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/drivermanagementmodule"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'drivermanagementmodule'} />
                </div>
              </PrivateRoute>
            }
          />
          <Route
            path="/notes/paymentmanagementmodule"
            element={
              <PrivateRoute>
                <div className="component">
                  <NotesList notesApp={'paymentmanagementmodule'} />
                </div>
              </PrivateRoute>
            }
          />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
