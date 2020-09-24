import React, { useState, useEffect } from "react";
import UploadService from "../services/FileUpload";
import Dashboard from "./Dashboard";
import TopItems from "./TopProfitableItems";
import { FaSpinner } from 'react-icons/fa';

function Uploader() {
    const [selectedFiles, setSelectedFiles] = useState(undefined);
    const [currentFile, setCurrentFile] = useState(undefined);
    const [progress, setProgress] = useState(0);
    const [message, setMessage] = useState("");
    const [records, setRecords] = useState([]);
    const [profit, setProfit] = useState([]);
    const [loading, setLoading] = useState(false)
    const [loading2, setLoading2] = useState(false)
    // useEffect(() => {
    //     UploadService.getFiles().then((response) => {
    //       setRecords(response.data);
    //     });
    //   }, []);

    const selectFile = (event) => {
        setSelectedFiles(event.target.files);
    };

    const upload = () => {
        let currentFile = selectedFiles[0];
    
        setProgress(0);
        setCurrentFile(currentFile);
    
        UploadService.upload(currentFile, (event) => {
          setProgress(Math.round((100 * event.loaded) / event.total));
        })
          .then((response) => {
            setLoading(true)
            setMessage(response.data);
            console.log(response.data)
            return UploadService.getRecords();
          })
          .then((records) => {
            setRecords(records.data);
            setLoading(false)
            setLoading2(true)
            console.log(records.data)
            return UploadService.topItems("2005-09-09", "2016-09-09")
          }).then((profit) => {
            console.log(profit.data)
            setProfit(profit.data)
            setLoading2(false)
            return UploadService.profitByDate("2005-09-09", "2016-09-09")
          }).then((profitDate) => {
            setMessage(profitDate.data)
          })
          .catch(() => {
            setProgress(0);
            setMessage("Could not upload the file!");
            setCurrentFile(undefined);
          });
    
        setSelectedFiles(undefined);
      };
  
    return (
        <div>
        {currentFile && (
          <div className="progress">
            <div
              className="progress-bar progress-bar-info progress-bar-striped"
              role="progressbar"
              aria-valuenow={progress}
              aria-valuemin="0"
              aria-valuemax="100"
              style={{ width: progress + "%" }}
            >
              {progress}%
            </div>
          </div>
        )}

        <label className="btn btn-default">
            <input type="file" onChange={selectFile} />
        </label>

      <button
        className="btn btn-success"
        disabled={!selectedFiles}
        onClick={upload}
      >
        Upload
      </button>

      {loading ? <FaSpinner className="fa-spin" /> : profit.length > 0 && <TopItems items={profit} />}
      {loading2 ? <FaSpinner className="fa-spin" /> : records.length > 0 && <Dashboard records={records} profit={message} />}

    </div>
    );
  }

export default Uploader;