import React, { useState, useEffect, useRef } from 'react';
import { Input, Message, Dropdown } from 'adesign-react';
import { Right as SvgRight, Down as SvgDown, Save as SvgSave } from 'adesign-react/icons';
import ApiStatus from '@components/ApiStatus';
import APIModal from '@components/ApisDescription';
import ManageGroup from '@components/ManageGroup';
// import { TYPE_MODAL_TYPE } from './types';
import Bus from '@utils/eventBus';
import './index.less';
import { useSelector, useDispatch } from 'react-redux';
import { fetchHandleApi } from '@services/apis';
import { cloneDeep } from 'lodash';
import { tap } from 'rxjs';
import { global$ } from '@hooks/useGlobal/global';
import { useTranslation } from 'react-i18next';
import Mousetrap from 'mousetrap';
import 'mousetrap-global-bind';
import useFolders from '@hooks/useFolders';
import { DropWrapper } from './style';
import { Button } from '@arco-design/web-react';
import { IconDown } from '@arco-design/web-react/icon';
import InputText from '@components/InputText';
import EnvView from '@components/EnvView';
const ButtonGroup = Button.Group;

const ApiInfoPanel = (props) => {
    const { data, showGenetateCode, onChange, from = 'apis', onSave } = props;


    const refDropdown = useRef();
    const [modalType, setModalType] = useState('');
    const open_apis = useSelector((store) => store.opens.open_apis);
    const open_api_now = useSelector((store) => store.opens.open_api_now);

    const { t } = useTranslation();
    const id_apis_scene = useSelector((store) => store.scene.id_apis);
    const id_now_scene = useSelector((store) => store.scene.id_now);

    const id_apis_plan = useSelector((store) => store.plan.id_apis);
    const id_now_plan = useSelector((store) => store.plan.id_now);

    const { apiFolders } = useFolders();

    const apiDataList = {
        'apis': open_apis,
        'scene': id_apis_scene,
        'plan': id_apis_plan,
    };
    const apiNowList = {
        'apis': open_api_now,
        'scene': id_now_scene,
        'plan': id_now_plan,
    };

    const apiData = apiDataList[from];
    const apiNow = apiNowList[from];

    const dispatch = useDispatch();

    const _saveId = useSelector((store) => store.opens.saveId);


    Mousetrap.bindGlobal(['command+s', 'ctrl+s'], () => {
        if (location.hash.split('/')[1] === 'apis') {
            saveApi(data.parent_id ? data.parent_id : '');
        }
        return false;
    });



    // const keyDown = (e) => {
    //     if (e.keyCode == 83 && (navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey)){
    //         e.preventDefault();
    //         saveApi();
    //      }
    // }

    // useEffect(() => {

    //     document.addEventListener('keydown', keyDown);

    //     return () => {
    //         document.removeEventListener('keydown', keyDown);
    //     }
    // }, []);

    const saveApi = (pid) => {
        if (data.name.trim().length === 0) {
            Message('error', t('message.apiNameEmpty'));
            return;
        }
        if (from === 'scene' || from === 'plan') {
            Bus.$emit('saveSceneApi', apiNow, apiData)
        } else {
            Bus.$emit('saveTargetById', {
                id: apiNow,
                saveId: _saveId,
                pid: pid ? pid : '0'
            }, {}, (code, id) => {

                if (code === 0) {
                    Message('success', t('message.saveSuccess'));

                    dispatch({
                        type: 'opens/updateSaveId',
                        payload: id,
                    })
                    // dispatch({
                    //     type: 'opens/updateOpenApiNow',
                    //     payload: id,
                    // })
                }
            });
        }
    }


    return (
        <>
            {modalType === 'description' && (
                <APIModal
                    value={data?.request?.description || ''}
                    onChange={onChange}
                    onCancel={setModalType.bind(null, '')}
                />
            )}
            <div className="api-manage">
                <div className="api-info-panel">
                    <div className="api-name-group">
                        <InputText
                            maxLength={30}
                            value={data.name || ''}
                            placeholder={t('placeholder.apiName')}
                            onChange={(e) => {
                                if (e.trim().length === 0) {
                                    Message('error', t('message.apiNameEmpty'));
                                    return;
                                }
                                onChange('name', e);
                            }}
                        />
                    </div>
                </div>

                <div className='info-panel-right'>
                    <EnvView env_id={data ? (data.env_info ? data.env_info.env_id : 0) : 0} onChange={onChange} />

                    <ButtonGroup>
                        <Button className="save-btn" onClick={() => saveApi(data.parent_id ? data.parent_id : '')}><SvgSave /> {t('btn.save')}</Button>

                        <Dropdown
                            ref={refDropdown}
                            placement="bottom-end"
                            content={
                                <div className={DropWrapper}>
                                    <div
                                        className="drop-item"
                                        onClick={(e) => {
                                            e.stopPropagation();
                                            saveApi();
                                            refDropdown.current.setPopupVisible(false);
                                        }}
                                    >
                                        {t('apis.rootFolder')}
                                    </div>
                                    {apiFolders.map((item) => (
                                        <>
                                            <div
                                                className="drop-item"
                                                key={item.target_id}
                                                {...item}
                                                value={item.target_id}
                                                onClick={() => {
                                                    console.log(item.target_id);
                                                    saveApi(item.target_id);
                                                    refDropdown.current.setPopupVisible(false);
                                                }}
                                            >
                                                {`|${new Array(item.level).fill('—').join('')}${item.name}`}
                                            </div>
                                        </>
                                    ))}
                                </div>
                            }

                        >
                            <Button className="save-more-btn" icon={<IconDown />} />

                        </Dropdown>
                    </ButtonGroup>
                    {/* <Button className='save-btn' onClick={() => saveApi()}
                    >{t('btn.save')}</Button> */}

                </div>
            </div>
        </>
    );
};

export default ApiInfoPanel;
