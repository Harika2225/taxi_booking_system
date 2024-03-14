import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import ClientHomePage from './components/pages/ClientHomePage/ClientHomePage';
import { Navigation } from './components/pages/Navigation/Navigation';

function App() {
  return (
    <div className="App">
      <Navigation />
      <Router>
        <Routes>
          <Route path="/" element={<ClientHomePage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
