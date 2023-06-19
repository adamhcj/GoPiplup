import './App.css';
import Gbtb from './Gbtb.js';
import Nparks from './Nparks.js';
import Select from 'react-select'
import { useState } from 'react';

function App() {
  const [place, setPlace] = useState('GBTB')

  return (
    <div>
      <div style={{'padding': '12px'}}>
        <h1>Welcome to GoPiplup</h1>
        <h3>Where would you like to Go?</h3>
        <h4 style={{'margin-bottom':'0px'}}>Please select a place to view events:</h4>
      </div>

      <div className='select-place-div'>
      <Select 
      options={[
        {value: 'GBTB', label: 'Gardens by the Bay'},
        {value: 'Nparks', label: 'Nparks'},
      ]}
      defaultValue={{value: 'GBTB', label: 'Gardens by the Bay'}}
      onChange={(e) => setPlace(e.value)}
      inputId='searchCompany'
      styles={{ menu: provided => ({ ...provided, zIndex: 9999 ,fontSize: '12px'}),
                control: provided => ({ ...provided, fontSize: '12px' }),
              }}
      />
      </div>
      
      {/* show depending on place */}
      <div style={{display: place === 'GBTB' ? 'block' : 'none'}}>
      <Gbtb />
      </div>
      <div style={{display: place === 'Nparks' ? 'block' : 'none'}}>
      <Nparks />
      </div>
  
    </div>
  );
}

export default App;
