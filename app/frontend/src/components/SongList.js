import React, { useState, useEffect } from 'react';
import axios from 'axios';

const SongList = () => {
    const [songs, setSongs] = useState([]);

    useEffect(() => {
        const fetchSongs = async () => {
            const result = await axios.get(process.env.REACT_APP_BACKEND_URL + '/api/songs');
            setSongs(result.data);
        };

        fetchSongs();
    }, []);

    return (
        <div>
            <h1>Song List</h1>
            <ul>
                {songs.map(song => (
                    <li key={song.id}>{song.title} by {song.artist}</li>
                ))}
            </ul>
        </div>
    );
};

export default SongList;
