import './App.css';
import Header from "./Components/header"
import Title from "./Components/title"
import Footer from "./Components/footer"

const App = () => {
  return (
    <div className="App">
      <Header />
      <Title />
      <div className="container">
      </div>
      <Footer />
    </div>
  );
}

export default App;
