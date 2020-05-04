import React from 'react';
import { ReactComponent as Chevron } from './chevron.svg'

import './styles.css'

export const BreadCrums = ({ history, breadcrums }) => {
    return <div className="breadcrums">
        {breadcrums.map((breadcrum, index) =>
            <>
                <div onClick={() => history.replace(breadcrum.url)}>
                    {breadcrum.label}
                </div>
              &nbsp;
              {index !== breadcrums.length - 1 &&
                    <>
                        <Chevron />
                  &nbsp;
                  </>
                }
            </>
        )}
    </div>
}