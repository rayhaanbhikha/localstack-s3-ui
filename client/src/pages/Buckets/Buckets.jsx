import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { WrapTable } from '../../Components'

const BucketRow = ({ bucketName }) => <tr>
    <td>
        <div>
            <Link to={`/s3/${bucketName}`}>
                {bucketName}
            </Link>
        </div>
    </td>
</tr>

export const Buckets = () => {
    const [data, setdata] = useState({})


    const fetchBuckets = async () => {
        try {
            console.log("loading bucket data")
            const res = await fetch(`http://localhost:8080/resource?path=/`)
            const data = await res.json();
            setdata(data.children)
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        fetchBuckets();
    }, [])

    const bucketNames = Object.entries(data).map(([_, { name }]) => name)

    const TableText = () => <>
        <strong className="table-bucket-text">Buckets</strong>
        &nbsp;&nbsp;
        <strong className="table-bucket-nums">({bucketNames.length})</strong>
    </>

    const TableBody = () =>
        bucketNames.map((bucketName, index) =>
            <BucketRow key={`bucketName-${index}`} bucketName={bucketName} />
        );

    return WrapTable(TableText, TableBody)
}
