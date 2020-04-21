import React, { useContext } from 'react'
import { S3Context } from '../../context'
import './styles.css'

export const Bucket = () => {
    const data = useContext(S3Context)

    return (
        <div className="buckets-table">
            <div className="table-head-container">
                <div className="table-text">
                    <strong className="table-bucket-text">Buckets</strong>&nbsp;&nbsp;<strong className="table-bucket-nums">()</strong>
                </div>
            </div>
            <table>
                <thead >
                    <tr className="table-column-heading">
                        <th className="table-column-heading-text">Name</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                    <tr>
                        <td>Some bucket</td>
                    </tr>
                </tbody>
            </table>
        </div>
    )
}
