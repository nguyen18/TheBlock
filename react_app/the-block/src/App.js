import './App.css';
import logo from './images/placeholder.png';

function App() {
  return (
    //<div className='AspectWrapper'>
      <div className='Content'>
        <h1 id='Welcome'>WELCOME TO</h1>
        <h1 id='The'>THE</h1>
        <img id='LandingLogo' src={logo} alt='the block landing logo'></img>
        <h1 id="Block">BLOCK</h1>
        <EnterButton name="Enter"></EnterButton>
      </div>
    //</div>

    // <div className='AspectWrapper'>
    //   <div className='Content'>
    //     <p>testing</p>
    //   </div>
    // </div>
      // <header className="App-header">
      //   <img src={logo} className="App-logo" alt="logo" />
      //   <p>
      //     Edit <code>src/App.js</code> and save to reload.
      //   </p>
      //   <a
      //     className="App-link"
      //     href="https://reactjs.org"
      //     target="_blank"
      //     rel="noopener noreferrer"
      //   >
      //     Learn React
      //   </a>
      // </header>
  );
}


const EnterButton = (props) => {
  return (
    <button type="button" id="EnterButton" href={props.link} target={props.target}>
      {props.name}
    </button>
  );
}

export default App;
