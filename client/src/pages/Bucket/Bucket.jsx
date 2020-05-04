import React, { useContext, useState, useEffect } from 'react';
import { withRouter } from 'react-router-dom'
import { ResourceRow, BreadCrums, Resource, WrapTable } from '../../Components'
import { S3Context } from '../../context'

import './styles.css'


const Component = ({ match }) => {

    const bucketName = match.params.bucketName;

    const [path, setPath] = useState("/" + bucketName)
    const [resources, setResources] = useState([])
    const [breadcrums, setBreadcrums] = useState([
        { label: "Localstack S3", url: "/s3" },
        { label: bucketName, url: `/s3/${path}` }
    ])

    const fetchResources = async () => {
        try {
            console.log("loading bucket data")
            const res = await fetch(`http://localhost:8080/resource?path=${path}`)
            const data = await res.json();
            setResources(data.children)
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        fetchResources();
    }, [path])

    const TableText = () => <>
        <BreadCrums breadcrums={breadcrums} />
        <strong className="table-bucket-text">{bucketName}</strong>
    </>

    const TableBody = () => resources.map((bucketData, index) =>
        <ResourceRow resource={bucketData} key={`${bucketData.name}-${index}`} setPath={setPath} />
    )

    return WrapTable(TableText, TableBody);
}

export const Bucket = withRouter(Component)