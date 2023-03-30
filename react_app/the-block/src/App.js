import './App.css';
import logo from './images/placeholder.png';

function App() {
  return (
    //<div className='AspectWrapper'>
      <div class='Content'>
        <div class="container">
          <div class='aligned'>
            <h1 id='Welcome'>WELCOME TO</h1>
          </div>
          <div class="blocktext">
            <div class='aligned'>
              <img id='LandingLogo' src={logo} alt='the block landing logo'></img>
            </div>
            <h1 id='top-left'>THE</h1>
            <h1 id="bottom-right">BLOCK</h1>
          </div>
          <div class = 'aligned'>
            <EnterButton name="Enter"></EnterButton>
          </div>
        </div>
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
