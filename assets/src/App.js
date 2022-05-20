import './App.css';
import Header from "./Components/header"
import Title from "./Components/title"
import Footer from "./Components/footer"
import TextAndButton from "./Components/textAndButton"
import TextArea from './Components/textArea';

const App = () => {
  return (
    <div className="App">
      <Header />
      <Title />
      <div className="container">
        <TextAndButton />
        <TextArea />
      </div>
      <Footer />
    </div>
  );
}

export default App;
