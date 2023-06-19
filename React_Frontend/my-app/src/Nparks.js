import React, { useEffect, useState } from 'react';

function Nparks() {

    const API_URL = process.env.REACT_APP_API_URL;

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        
        
        fetch(`${API_URL}/nparks`)
        .then(response => response.text())
        .then(data => {
            setLoading(false);
            document.getElementById("nparks-main-div").innerHTML = data;
        }
        )
        .catch(err => {
            setLoading(false);
            document.getElementById("nparks-main-div").innerHTML = err;
        });





    }, []);



  return (
    <div className='gbtb-container'>
        {loading && 
        <div class="loader">
            <div class="loader-wheel"></div>
            <div class="loader-text"></div>
        </div>
        }

    <div id="nparks-main-div">
    </div>
        
    </div>
  );
}

export default Nparks;