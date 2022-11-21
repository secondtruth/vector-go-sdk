package sdk_wrapper

import (
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

/* Animationnames on Vector 1.8
anim_reacttocliff_stuckonedge_alert_left_01
anim_slowpoke_loop_04
anim_observing_sides_area_01
anim_message_deleted_short_01
anim_heldonpalm_putdown_relaxed_01
anim_fastbump_loop_01
anim_hiking_driving_start_04
anim_eyecontact_getout_01_head_angle_-20
anim_lookinplaceforfaces_keepalive_long_03
anim_volume_stage_01
anim_keepaway_pounce_shake_02
anim_keepalive_eyes_01_updown
anim_tts_loop_02
anim_blackjack_speech_tts_01
anim_eyecontact_left_contact_01_head_angle_20
anim_reacttoblock_react_01_head_angle_40
anim_blackjack_victorbust_01
anim_vc_reaction_nofaceheardyou_01
anim_reacttocliff_stuckonedge_alert_right_01
anim_reacttocliff_turnleft_60_01
anim_observe_oncharger_getout_01
anim_explorer_planning_getin_02
anim_eyepose_sad_instronspect
anim_launch_reacttocube
anim_wakeword_getout_turnright_90degrees_01
anim_howold_getin_even_01
anim_reacttoblock_ask_01_0
anim_keepaway_hit_reaction_01
anim_pause_idle_03
anim_gazing_lookatsurfaces_getin_left_01_head_angle_20
anim_explorer_scan_short_01
anim_rtsound_oncharger_observe_behind_left_01
anim_explorer_scan_short_03
anim_explorer_scan_short_02
anim_launch_wakeup_startdriving_01
anim_explorer_scan_short_04
anim_explorer_idle_02_head_angle_-10
anim_reacttocliff_wheely_01
anim_reacttoblock_lifteffortpickup_01
anim_chargerdocking_requestgetout_01
anim_lookinplaceforfaces_keepalive_long_04
anim_keepalive_eyes_01_forward
anim_reacttocliff_edge_04
anim_snowglobe_loop_01
anim_rtsound_offcharger_asleep_60left_01
anim_knowledgegraph_searching_01
anim_dancebeat_headliftbody_left_med_01
anim_petting_lvl3_getout_02
anim_hiking_observe_01
anim_wakeword_getout_groggyeyes_turnright_135degrees_01
anim_reacttoface_unidentified_03_head_angle_40
anim_chargerdocking_liftup_01
anim_explorer_idle_03_head_angle_40
anim_cubedocking_drive_getout_01
anim_lookinplaceforfaces_keepalive_long_02
anim_referencing_squint_01_head_angle_40
anim_freeplay_reacttoface_longname_01
anim_fistbump_requesttwice_01
anim_weather_windy_01
anim_movement_right_01
anim_heldonpalm_edge_nervous_01
anim_cubeconnection_getin_01
anim_reacttocliff_sidestuck_getout_01
anim_gazing_lookatvector_reaction_01_head_angle_20
anim_handdetection_reaction_02
anim_referencing_curious_01_head_angle_20
anim_eyecontact_center_contact_01_head_angle_-20
anim_rtsound_oncharger_asleep_60right_01
anim_rtsound_offcharger_observe_left_01
anim_lookatdevice_getin
anim_blackjack_victorbjacklose_01
anim_eyepose_sad
anim_eyecontact_smile_01_head_angle_20
anim_rtsound_oncharger_observe_30left_01
anim_rtpickup_reaction_02
anim_cubedocking_drive_getin_01
anim_rtpickup_reaction_01
anim_knowledgegraph_listening_01
anim_reacttoface_unidentified_01
anim_reacttoface_unidentified_02
anim_vc_laser_lookdown_01
anim_vc_listening_01
anim_photo_focus_01
anim_keepaway_winhand_03
anim_referencing_giggle_01_head_angle_-20
anim_lookatdevice_icon
anim_lookinplaceforfaces_keepalive_medium_head_angle_40
anim_keepalive_eyes_01_right
anim_uperformance_loosepixel_01
anim_pounce_fail_02
anim_heldonpalm_transition2relaxed_01
anim_gif_idk_01
anim_chargerdocking_leftturn_all
anim_gazing_lookatfaces_turn_left_01
anim_movement_alreadyhere_01_head_angle_-20
anim_handdetection_getout_01
anim_dizzy_reaction_medium_02
anim_chargerdocking_comeoff_right_02
anim_dancebeat_cantdothat_01
anim_chargerdocking_comeoff_right_01
anim_lookinplaceforfaces_keepalive_long_02_head_angle_-20
anim_chargerdocking_searchforcharger_canthelp
anim_vc_listening_loop_02
anim_pounce_lookloop_03
anim_keepaway_getout_exit_01
anim_eyepose_determined
anim_eyecontact_lookloop_02
anim_findcube_searchforcubesturns_02
anim_findcube_searchforcubesturns_01
anim_freeplay_reacttoface_sayname_03_head_angle_-20
anim_qa_motors_90leftbackward_arcturn_500ms_01
anim_message_play_reaction_01
anim_eyepose_captivated
anim_power_offon_02
anim_power_offon_04
anim_eyepose_angry
anim_lookinplaceforfaces_keepalive_short_head_angle_20
anim_rtshake_drive_loop_02
anim_reacttoface_unidentified_03
anim_movement_lookinplaceforfaces_long
anim_onboarding_drive_to_face_01
anim_petdetection_reaction_dog_03
anim_qa_stub
anim_eyecontact_right_contact_01_head_angle_-20
anim_holiday_hh_lights_01
anim_holiday_hh_lights_02
anim_gotosleep_getout_01
anim_gotosleep_getout_03
anim_cubespinner_sessionfail_01
anim_rtsound_oncharger_observe_30right_01
anim_reacttocliff_wheely_03
anim_reacttocliff_stuckonedge_02
anim_onboarding_driveoff_charger_alt_01
anim_knowledgegraph_getout_01
anim_dancebeat_headliftbody_back_01
anim_blackjack_idle_01
anim_qa_sync_5sec
anim_rtpkeepaway_idle_02
anim_weather_rain_01
anim_explorer_huh_01_head_angle_40
anim_cubedocking_fail_01
anim_frequency_test_02
anim_frequency_test_01
anim_frequency_test_04
anim_weather_cloud_01
anim_rtsound_oncharger_observe_left_01
anim_chargerdocking_severergetout_02
anim_pounce_lookloop_02
anim_reacttoblock_dropfail_02
anim_eyepose_asleep
anim_generic_look_stright_01
anim_reacttocliff_wheely_02
anim_gazing_lookatsurface_reaction_02
anim_hiking_driving_end_01
anim_gotosleep_getout_04
anim_fistbump_requestoncelong_01
anim_reacttocliff_faceplantroll_01
anim_slowbump_loop_01
anim_reacttoface_unidentified_02_head_angle_20
anim_rtpickup_loop_03
anim_cubespinner_faketap_01
anim_observing_nearby_area_02
anim_rtpickup_loop_05
anim_rtpickup_loop_09
anim_rtillumination_lightonwakeup_01
anim_petting_blissloop_03
anim_dancebeat_headliftbody_left_small_01
anim_gazing_lookatvector_reaction_01
anim_react2block_02
anim_reacttocliff_stuckrightside_01
anim_reacttocliff_stuckrightside_03
anim_fistbump_getin_01
anim_reacttocliff_stuckrightside_02
anim_onboarding_wakeword_getout_02
anim_rtshake_lv2pregetout_01
anim_howold_getout_01
anim_freeplay_reacttoface_sayname_02_head_angle_40
anim_lookinplaceforfaces_keepalive_long_04_head_angle_40
anim_reacttocliff_pickup_01
anim_power_offon_01
anim_reacttocliff_pickup_03
anim_reacttocliff_pickup_05
anim_lookinplaceforfaces_keepalive_medium
anim_blackjack_getin_01
anim_wakeword_groggyeyes_getout_01
anim_rtshake_lv2rtonground_looseeye_01
anim_petting_blissloop_01
anim_reacttocliff_edgeliftup_01
anim_dancebeat_headliftbody_left_large_01
anim_rtsound_oncharger_asleep_front_01
anim_keepaway_idle_side_02
anim_explorer_huh_01_head_angle_-20
anim_look_left_01
anim_rtsound_offcharger_observe_120right_01
anim_eyepose_normal_l_push
anim_qa_motors_90rightbackward_arcturn_500ms_01
anim_referencing_smile_01_head_angle_20
anim_vc_alrighty_01
anim_eyecontact_right_thought_01_head_angle_-20
anim_meetvictor_lookface_interrupted_01
anim_rtp_blackjack_idle_01
anim_meetvictor_lookface_loop_03
anim_referencing_curious_01
anim_reacttocliff_faceplantroll_armup_01
anim_hiking_driving_start_03
anim_inspectheldcube_getin_01
anim_rtsound_offcharger_asleep_ambient_01
anim_rtshake_lv3rtonhand_01
anim_cubespinner_loop_03
anim_cubespinner_loop_02
anim_cubeconnection_connectionlost_01
anim_fistbump_success_01
anim_fistbump_success_03
anim_eyepose_shocked
anim_rtmotion_up_01
anim_reacttoblock_success_01
anim_gazing_lookatsurfaces_getin_right_01_head_angle_20
anim_eyecontact_center_contact_01_head_angle_40
anim_movement_comehere_greeting_02
anim_timerleft_getout_01
anim_eyepose_unsure
anim_pounce_success_03
anim_pounce_success_01
anim_rtshake_lv2rtonhand_01
anim_wakeword_getout_groggyeyes_turnright_45degrees_01
anim_cubespinner_gamesuccess_01
anim_movement_alreadyhere_01_head_angle_40
anim_meetvictor_alreadyknowbob_01
anim_chargerdocking_comeoff_straight_02
anim_rtsound_offcharger_asleep_150right_01
anim_lookinplaceforfaces_keepalive_long_03_head_angle_20
anim_lookinplaceforfaces_keepalive_short
anim_explorer_huh_close_01
anim_keepaway_pounce_slow_05
anim_keepaway_driving_loop_01
anim_onboarding_wakeword_loop_02
anim_keepaway_driving_loop_02
anim_chargerdocking_searchforchargerturns_default_03
anim_onboarding_lookdown_lookaround_loop_01
anim_nowifi_signal_01
anim_rtshake_lv3loop_01
anim_feedback_badrobot_02
anim_communication_cantdothat_04
anim_rtsound_oncharger_asleep_150left_01
anim_qa_motors_lift_down_500ms_01
anim_eyecontact_left_thought_01_head_angle_20
anim_gotosleep_wakeup_01
anim_message_announce_toplayer_01
anim_chargerdocking_searchforchargerturns_default_loop_01
anim_keepaway_getin_03
anim_rtsound_offcharger_asleep_150left_01
anim_knowledgegraph_fail_getout_01
anim_reacttoblock_focusedeyes_01
anim_fistbump_requestonce_01
anim_pause_idle_01
anim_keepaway_fakeout_06
anim_reacttocliff_reaction_front_right_01
anim_movement_alreadyhere_01
anim_wakeword_loop_01
anim_greeting_goodmorning_02
anim_lookinplaceforfaces_keepalive_long_02_head_angle_40
anim_heldonpalm_pickup_nervous_01
anim_fastbump_getin_01
anim_chargerdocking_severergetout_01
anim_freeplay_reacttoface_sayname_03_head_angle_20
anim_gazing_lookatsurfaces_turn_right_01
anim_test_spriteboxremaps
anim_eyecontact_center_contact_01_head_angle_20
anim_movement_reacttoface_01_head_angle_-20
anim_eyepose_blink_low
anim_qa_sync_lift
anim_lowlightcharger_search_getout_01
anim_feedback_shutup_02
anim_reacttoblock_frustrated_01
anim_qa_head_updown_alt
anim_blackjack_ok_01
anim_inspectheldcube_loop_03
anim_keepaway_getin_focus_01
anim_keepaway_getout_loseinterest_01
anim_inspectheldcube_loop_01
anim_gotosleep_getin_01
anim_reacttocliff_stuckonedge_left_02
anim_explorer_huh_01_head_angle_20
anim_movement_comehere_01
anim_reacttoface_unidentified_03_head_angle_20
anim_onboarding_cube_reacttocube
anim_gazing_lookatfaces_getin_right_01_head_angle_40
anim_feedback_apology_01
anim_eyecontact_right_contact_01
anim_gazing_lookatsurface_reaction_01
anim_explorer_idle_03_head_angle_10
anim_triple_backup
anim_rtshake_lv3rtonground_01
anim_feedback_apology_02
anim_meetvictor_sayname_03
anim_freeplay_reacttoface_sayname_01_head_angle_40
anim_explorer_scan_right_01
anim_wakeword_getout_turnleft_180degrees_01
anim_explorer_scan_right_02
anim_greeting_goodnight_02
anim_chargerdocking_searchforchargerturns_default_02
anim_dizzy_reaction_soft_03
anim_movement_turnaround_01
anim_lookinplaceforfaces_keepalive_long_head_angle_-20
anim_dancebeat_scoot_right_01
anim_fistbump_idle_02
anim_chargerdocking_comeoff_binaryeyes_01
anim_wakeword_getin_01
anim_lowlightcharger_search_loop_02
anim_cubeconnection_connectionsuccess_01
anim_rtshake_lv2rtonhand_looseeye_01
anim_slowbump_getin_01
anim_reacttoblock_react_01_head_angle_-20
anim_rtpmemorymatch_request_03
anim_reacttocliff_faceplantroll_02
anim_eyepose_suspicious
anim_reacttoblock_happydetermined_01
anim_keepaway_losegame_01
anim_launch_drivingloop_01
anim_keepaway_idle_glance_01
anim_launch_drivingloop_03
anim_launch_drivingloop_02
anim_cubespinner_loop_04
anim_freeplay_reacttoface_identified_01
anim_freeplay_reacttoface_identified_02
anim_reacttoblock_reacttocharger_01
anim_wakeword_turn_left_01
anim_lowlightcharger_search_getin_02
anim_qa_sync_head
anim_lowlightcharger_search_getin_01
anim_referencing_squint_01_head_angle_20
anim_chargerdocking_searchforchargerturns_default_loop_02
anim_fistbump_fail_01
anim_reacttocliff_sidestuck_effort_01
anim_driving_upset_loop_01
anim_knowledgegraph_fail_01
anim_movement_alreadyhere_01_head_angle_20
anim_nowifi_getin_02
anim_eyepose_blink_bd_l
anim_eyepose_all_faces
anim_getin_tts_01
anim_cubespinner_rtcubemoved_01
anim_keepaway_fakeout_01
anim_greeting_goodbye_01
anim_vc_reaction_nofaceheardyou_getout_01
anim_communication_cantdothat_02
anim_slowpoke_getout_04
anim_reacttoblock_dropfail_01
anim_onboarding_reacttoface_happy_01_head_angle_20
anim_rtpickup_loop_01
anim_sdk_speak_01
anim_greeting_goodbye_02
anim_referencing_curious_01_head_angle_40
anim_volume_stage_03
anim_lookatdvice_getout
anim_keepaway_getin_01
anim_eyepose_curious
anim_nowifi_getin_01
anim_wakeword_getout_01
anim_rtpickup_loop_06
anim_tts_loop_03
anim_volume_stage_max
anim_rtsound_offcharger_asleep_behind_right_01
anim_reacttocliff_stuckleftside_01
anim_reacttocliff_stuckonedge_left_getin_01
anim_reacttoblock_react_short_01
anim_observing_around_subtle_02
anim_dancebeat_getin_02
anim_observing_nearby_area_01
anim_timerset_getin_01
anim_rtmotion_45left_01
anim_eyepose_sad_up
anim_qa_motors_head_up_500ms_01
anim_gazing_lookatfaces_getin_left_01_head_angle_40
anim_onboarding_wakeword_getin_01
anim_rtsound_oncharger_observe_behind_right_01
anim_qa_firmwaremessaging_01
anim_rtsound_oncharger_observe_150right_01
anim_petting_lvl4_01
anim_cubeconnection_loop_03
anim_codelab_getout_01
anim_turn_to_motion_pounce_right_01
anim_launch_wakeup_01
anim_message_tts
anim_launch_wakeup_04
anim_gazing_lookatfaces_turn_right_01
anim_launch_wakeup_03
anim_launch_wakeup_05
anim_test_spriteboxkeyframes
anim_onboarding_wakeup_01
anim_launch_wakeup_drivingloop_02
anim_keepaway_pounce_getin_01
anim_handdetection_lift_01
anim_launch_cubediscovery
anim_wakeword_getout_groggyeyes_turnleft_135degrees_01
anim_keepalive_eyesonly_loop_02
anim_reacttoface_unidentified_02_head_angle_40
anim_petdetection_reaction_cat_02
anim_eyecontact_lookloop_02_head_angle_-20
anim_keepalive_eyesonly_loop_01
anim_hiking_driving_loop_05
anim_reacttoblock_lifteffortplacehigh_01
anim_explorer_planning_getout_02
anim_gazing_lookatsurfaces_getin_left_01
anim_meetvictor_sayname02_04
anim_pounce_fail_01
anim_eyepose_happy
anim_explorer_huh_far_01
anim_dancebeat_eyebeat_01
anim_referencing_smile_01_head_angle_-20
anim_wakeword_idk_02
anim_explorer_scan_left_02
anim_rtshake_lv1pregetout_01
anim_explorer_scan_left_01
anim_eyecontact_giggle_01_head_angle_20
anim_keepaway_idle_inplace_03
anim_hiking_react_05
anim_gotosleep_sleeploop_01
anim_freeplay_reacttoface_identified_03
anim_hiking_lookaround_02
anim_chargerdocking_searchforchargerturns_03
anim_rtsound_oncharger_asleep_150right_01
anim_explorer_idle_01
anim_micstate_micon_01
anim_rtpickup_loop_02
anim_loco_driving01_loop_03
anim_rtsound_offcharger_observe_right_01
anim_rtsound_offcharger_asleep_left_01
anim_rtpickup_putdown_03
anim_loco_driving01_loop_01
anim_rtmotion_lookleft_01
anim_referencing_squint_02_head_angle_40
anim_fastbump_getout_01
anim_chargerdocking_pickup_01
anim_howold_getin_odd_01
anim_reacttoblock_lifteffortroll_01
anim_eyepose_normal_r
anim_dancebeat_headliftbody_right_small_01
anim_referencing_squint_02_head_angle_20
anim_blackjack_goodluck_01
anim_exploring_intervals_test
anim_onboarding_eyecontact_loop_getout_01
anim_freeplay_reacttoface_sayname_02_head_angle_-20
anim_rtpickup_loop_07
anim_rtp_blackjack_timeout_01
anim_keepaway_getready_02
anim_heldonpalm_jolt_01
anim_referencing_smile_01
anim_chargerdocking_searchforchargerturns_default_getout
anim_eyecontact_getout_01_head_angle_40
anim_attention_lookatdevice_01
anim_observing_around_subtle_01
anim_reacttoface_unidentified_02_head_angle_-20
anim_pounce_success_04
anim_observing_far_subtle_01
anim_rtsound_offcharger_observe_150left_01
anim_hiking_driving_end_02
anim_eyepose_normal_r_push
anim_explorer_idle_02
anim_timersup_getout_01
anim_wakeword_turn_right_135degrees_01
anim_cubespinner_anticsessionfail_01
anim_rtsound_offcharger_observe_120left_01
anim_weather_thunderstorm_01
anim_cube_success_getout_01
anim_rtmotion_45right_notreads_01
anim_gazing_lookatsurfaces_getin_left_01_head_angle_40
anim_pounce_success_02
anim_qa_motors_90leftforward_arcturn_500ms_01
anim_blackjack_victorwin_01
anim_referencing_squint_01_head_angle_-20
anim_reacttoblock_react_short_02
anim_petting_blissloop_02
anim_wakeword_getout_turnright_135degrees_01
anim_eyepose_blink
anim_blackjack_deal_01
anim_timer_beep_02
anim_dizzy_reaction_soft_01
anim_fistbump_idle_03
anim_rtmotion_left_getout_01
anim_iperformance_tangram_01
anim_face_sleeping
anim_petdetection_reaction_cat_01
anim_eyecontact_smile_01_head_angle_40
anim_slowpoke_getin_02
anim_fistbump_idle_01
anim_movement_comehere_greeting_01
anim_hiking_react_04
anim_eyepose_joy
anim_petting_lvl2_01
anim_rtmotion_forwardup_01
anim_greeting_hello_01
anim_greeting_hello_02
anim_timerleft_getin_01
anim_movement_lookinplaceforfaces_medium_head_angle_-20
anim_lookinplaceforfaces_keepalive_long_head_angle_40
anim_weather_snow_01
anim_wakeword_groggyeyes_listenloop_01
anim_eyepose_focused
anim_chargerdocking_comeoff_left_01
anim_howold_fastloop_01
anim_rtpkeepaway_idle_01
anim_eyecontact_right_contact_01_head_angle_20
anim_reacttoface_unidentified_01_head_angle_20
anim_petting_lvl4_getout_01
anim_movement_getout_01
anim_keepaway_losehand_03
anim_reacttocliff_pickup_04
anim_rtmotion_turn45right_01
anim_referencing_giggle_01
anim_onboarding_lookaround_loop_01
anim_reacttoblock_react_01_head_angle_20
anim_rtshake_lv1rtonground_01
anim_eyecontact_left_thought_01_head_angle_40
anim_weather_sunny_01
anim_chargerdocking_settle_01
anim_timer_emote_01
anim_snowglobe_pregetout_01
anim_hiking_react_06
anim_feedback_bequiet_01
anim_eyepose_frustrated
anim_loco_driving01_start_01
anim_inspectheldcube_timeout_01
anim_dancebeat_headliftbody_fwd_01
anim_movement_lookinplaceforfaces_short_head_angle_40
anim_reacttocliff_stuckleftside_getin_01
anim_reacttocliff_stuckonedge_right_02
anim_movement_lookinplaceforfaces_long_head_angle_20
anim_onboarding_cube_success_getout_01
anim_eyecontact_right_thought_01_head_angle_20
anim_holiday_hny_fireworks_01
anim_wakeword_getout_turnleft_135degrees_01
anim_reacttocliff_turtlerollfail_03
anim_greeting_happy_03
anim_weather_thunderstorm_01_temp_remap
anim_knowledgegraph_searching_getin_01
anim_slowbump_getout_01
anim_vc_listening_getout_01
anim_referencing_squint_01
anim_cubespinner_reaction_01
anim_driving_upset_loop_03
anim_greeting_goodnight_01
anim_explorer_huh_01_head_angle_-10
anim_rtsound_oncharger_observe_60left_01
anim_pairing_icon_awaitingapp
anim_hiking_lookaround_03
anim_keepaway_bored_idle_01
anim_frequency_test_03
anim_keepaway_bored_idle_02
anim_cubespinner_anticsessionsuccess_01
anim_rtp_blackjack_playeryes_01
anim_power_onoff_01
anim_explorer_idle_01_head_angle_-20
anim_hiking_driving_loop_03
anim_keepaway_getout_frustrated_01
anim_rtsound_offcharger_observe_60left_01
anim_feedback_iloveyou_01
anim_reacttocliff_stuckleftside_03
anim_reacttocliff_turnleft_120_01
anim_tts_loop_01
anim_dizzy_shake_loop_01
anim_fistbump_requestonce_02
anim_dancebeat_pivot_left_01
anim_neutral_eyes_01
anim_lookinplaceforfaces_keepalive_long_03_head_angle_40
anim_chargerdocking_searchforchargerturns_default_01
anim_cube_drive_loop_01
anim_reacttoblock_stuckonblock_01
anim_explorer_huh_01
anim_dizzy_reaction_medium_01
anim_cubespinner_anticgamefail_01
anim_explorer_idle_01_head_angle_10
anim_freeplay_reacttoface_sayname_02
anim_dizzy_pickup_01
anim_referencing_giggle_01_head_angle_20
soundTestAnim
anim_snowglobe_rtonhand_01
anim_hiking_driving_loop_01
anim_pause_idle_02
anim_hiking_driving_loop_06
anim_meetvictor_lookface_loop_02
anim_blackjack_victorlose_01
anim_eyecontact_squint_01_head_angle_20
anim_eyecontact_lookloop_02_head_angle_40
anim_look_right_01
anim_heldonpalm_rolloff_01
anim_eyecontact_left_contact_01
anim_turn_right_01
anim_reacttocliff_stuckleftside_02
anim_wakeword_getout_turnleft_90degrees_01
anim_reacttocliff_turnleft_180_01
anim_hiking_react_07
anim_knowledgegraph_searching_getout_01
anim_loco_driving01_start_02
anim_movement_lookinplaceforfaces_short
anim_wakeword_idk_04
anim_eyepose_bliss
anim_eyecontact_smile_01
anim_powersavemode_temperature_01
anim_keepaway_losegame_02
anim_timercancel_02
anim_volume_stage_02
anim_explorer_idle_03
anim_meetvictor_getin_01_head_angle_40
anim_loco_driving01_end_01
anim_eyecontact_center_thought_01_head_angle_40
anim_spinner_tap_01
anim_reacttocliff_turtleroll_04
anim_rtsound_offcharger_asleep_120left_01
anim_freeplay_reacttoface_sayname_03_head_angle_40
anim_blackjack_victorbjackwin_01
anim_snowglobe_getin_01
anim_rtpmemorymatch_request_01
anim_rtpkeepaway_reacttocube_03
anim_rtpickup_loop_08
anim_onboarding_reacttoface_happy_01
anim_rtmotion_45left_notreads_01
anim_explorer_idle_03_head_angle_20
anim_inspectheldcube_loop_04
anim_qa_sync_treads
anim_qa_lift_updown
anim_keepaway_wingame_02
anim_explorer_center_right_01
anim_knowledgegraph_getin_01
anim_pairing_icon_update_error
anim_explorer_idle_03_head_angle_30
anim_dancebeat_getout_01
anim_explorer_idle_01_head_angle_-10
anim_observing_self_absorbed_01
anim_rtmotion_lookup_01
anim_reacttocliff_sidestuck_effort_02
anim_eyecolorreact_switch_01
anim_heldonpalm_getin_nervous_01
anim_petdetection_reaction_cat_03
anim_eyecolorreact_getout_01
anim_volume_stage_04
anim_handdetection_drive_loop_01
anim_keepaway_winhand_01
anim_eyecontact_center_thought_01
anim_reacttocliff_edge_02
anim_rtpmemorymatch_request_02
anim_hiking_driving_loop_04
anim_chargerdocking_severerequest_01
anim_eyecontact_left_thought_01
anim_reacttoblock_placeblock_01
anim_nocloud_icon_01
anim_meetvictor_lookface_timeout_01
anim_gazing_lookatsurfaces_getin_right_01
anim_keepaway_stopshort_01
anim_dancebeat_listening_02
anim_pounce_fail_03
anim_findcube_searchforcubesturns_03
anim_cube_psychic_01
anim_eyecontact_lookloop_01_head_angle_-20
anim_fistbump_success_02
anim_petting_lvl3_01
anim_blackjack_swipe_01_temp_remap
anim_wakeword_getout_turnright_45degrees_01
anim_rtsound_oncharger_asleep_120left_01
anim_rtsound_oncharger_asleep_right_01
anim_rtpkeepaway_reacttocube_01
anim_movement_reacttoface_01_head_angle_20
anim_reacttocliff_faceplantroll_armup_02
anim_feedback_shutup_01
anim_sudden_obstacle_react_02
anim_movement_left_02
anim_greeting_imhome_01
anim_findcube_request_01
anim_eyecolorreact_getin_01
anim_sudden_obstacle_react_01
anim_gazing_lookatfaces_getin_left_01_head_angle_-20
anim_launch_wakeup_02
soundTestAnimBell
anim_dancebeat_headnod_down_01
anim_onboarding_drive_getout_01
anim_cubespinner_anticroundsuccess_01
anim_explorer_idle_01_head_angle_20
anim_reacttocliff_turnright_120_01
anim_dancebeat_eyebeat_right_01
anim_chargerdocking_searchforcharger_01
anim_movement_lookinplaceforfaces_short_head_angle_-20
anim_rtmotion_45right_01
anim_onboarding_driveoff_charger_01
anim_wakeword_getout_turnleft_45degrees_01
anim_eyepose_blink_bd_r
anim_freeplay_reacttoface_like_01
anim_rtsound_oncharger_asleep_60left_01
anim_timer_beep_01
anim_eyepose_awe
anim_explorer_planning_getout_01
anim_rtsound_offcharger_observe_behind_left_01
anim_rtsound_offcharger_observe_150right_01
anim_weather_stars_01_temp_remap
anim_snowglobe_rtonground_01
anim_rtsound_oncharger_asleep_behind_left_01
anim_wakeword_getout_groggyeyes_turnleft_45degrees_01
anim_slowpoke_getinbackward_01
anim_handdetection_reaction_01
anim_generic_look_up_idle_02
anim_meetvictor_getin_01
anim_movement_reacttoface_01
anim_lookatdevice_icon_start
anim_reacttocliff_turtleroll_05
anim_petting_bliss_getout_02
anim_wakeword_groggyeyes_getin_01
anim_reacttocliff_stuckonedge_right_01
anim_lookinplaceforfaces_keepalive_long_03_head_angle_-20
anim_keepaway_pounce_mousetrap_04
anim_cubeconnection_loop_01
anim_petting_lvl2_getout_02
anim_meetvictor_sayname02_02
anim_communication_cantdothat_03
anim_rtsound_offcharger_observe_30right_01
anim_hiking_react_02
anim_weather_cold_01_temp_remap
anim_reacttocliff_pickup_02
anim_reacttocliff_turtleroll_02
anim_eyepose_normal_l
anim_eyepose_scared
anim_qa_sync_5_10_50
anim_onboarding_wakeword_loop_01
anim_petdetection_reaction_dog_02
anim_heldonpalm_pickup_relaxed_01
anim_reacttoface_unidentified_01_head_angle_-20
anim_rtsound_offcharger_observe_behind_right_01
anim_reacttocliff_turtleroll_03
anim_cubespinner_gamefail_01
anim_eyepose_startled
anim_rtsound_offcharger_observe_ambient_01
anim_observing_distance_area_01
anim_explorer_idle_02_head_angle_40
anim_freeplay_reacttoface_sayname_01_head_angle_-20
anim_reacttoblock_dropsuccess_01
anim_slowpoke_loop_01
anim_wakeword_getout_groggyeyes_turnleft_90degrees_01
anim_gazing_lookatvector_reaction_01_head_angle_-20
anim_eyecontact_lookloop_01_head_angle_40
anim_rtsound_offcharger_observe_60right_01
anim_handdetection_getin_01
anim_laser_drive_01
anim_rtsound_offcharger_asleep_right_01
anim_chargerdocking_request1_01
anim_dancebeat_pivot_right_01
anim_explorer_idle_01_head_angle_30
anim_keepaway_fakeout_03
anim_petting_lvl3_getout_01
anim_feedback_goodrobot_02
anim_qa_motors_lift_up_500ms_01
anim_hiking_react_03
anim_qa_motors_forward_500ms_01
anim_keepaway_pounce_quick_01
anim_heldonpalm_getin_relaxed_01
anim_petting_lvl1_getout_01
anim_hiking_react_01
anim_dancebeat_getready_01
anim_reacttocliff_reaction_rear_left_01
anim_generic_look_up_01
anim_explorer_idle_03_head_angle_-10
anim_eyecontact_left_thought_01_head_angle_-20
anim_communication_cantdothat_01
ANIMATION_TEST
anim_lookinplaceforfaces_keepalive_short_head_angle_40
anim_rtshake_drive_getin_01
anim_reacttocliff_edge_03
anim_micstate_micoff_01
anim_keepaway_getreadyset_01
anim_reacttocliff_turtleroll_01
anim_loco_driving01_loop_02
anim_gotosleep_sleeping_05
anim_movement_lookinplaceforfaces_short_head_angle_20
anim_freeplay_reacttoface_sayname_01_head_angle_20
anim_explorer_idle_03_head_angle_-20
anim_rtmotion_up_notreads_01
anim_launch_reacttoputdown
anim_timersup_beep_01
anim_rtpickup_putdown_02
anim_gazing_lookatfaces_getin_right_01_head_angle_-20
anim_rtmotion_up_getout_01
anim_rtsound_offcharger_asleep_60right_01
anim_timersup_getin_02
anim_timersup_getin_01
anim_pounce_02
anim_exploring_intervals_4sslows
anim_feedback_bequiet_02
anim_reacttoblock_admirecubetower_01
anim_pairing_icon_wifi
anim_rtshake_lv2loop_01
anim_rtshake_drive_loop_03
anim_timerset_getout_01
anim_reacttoblock_putdown2nd_01
anim_reacttocliff_stuckonedge_left_01
anim_keepalive_eyesonly_loop_04
anim_reacttocliff_stuckrightside_getin_01
anim_onboarding_reacttoface_happy_01_head_angle_-20
anim_eyecontact_center_thought_01_head_angle_20
anim_meetvictor_sayname02_01
anim_dancebeat_headnod_01
anim_keepalive_blink_01
anim_rtpickup_reaction_03
anim_hiking_driving_start_02
anim_meetvictor_getin_01_head_angle_-20
anim_heldonpalm_looking_nervous_01
anim_blackjack_quit_01
anim_turn_to_motion_pounce_left_01
anim_reacttoblock_lifteffortplacelow_01
anim_explorer_planning_idle_01
anim_lookinplaceforfaces_keepalive_medium_head_angle_-20
anim_lowlightcharger_search_loop_01
anim_explorer_idle_02_head_angle_-20
anim_rtshake_lv1rtonhand_01
anim_keepaway_wingame_01
anim_rtsound_oncharger_observe_front_01
anim_meetvictor_sayname_04
anim_inspectheldcube_cubemissing_01
anim_chargerdocking_rightturn_all
anim_keepaway_fakeout_04
anim_holiday_hyn_confetti_01
anim_inspectheldcube_cubeonground_01
anim_keepaway_stopshort_02
anim_eyecontact_getout_01
anim_eyepose_squint
anim_explorer_idle_02_head_angle_10
anim_chargerdocking_comeoff_left_02
anim_wakeword_getout_groggyeyes_turnleft_180degrees_01
anim_launch_sleeping_01
anim_qa_motors_90right_turn_500ms_01
anim_message_loop_01
anim_rtsound_oncharger_observe_ambient_01
anim_eyecolorreact_idle_01
anim_dancebeat_headliftbody_right_med_01
anim_dancebeat_eyebeat_left_01
anim_rtmotion_lookright_01
anim_communication_pickup_cantdothat_01
anim_keepaway_getready_03
anim_gazing_lookatfaces_getin_right_01
anim_eyecontact_right_thought_01_head_angle_40
anim_referencing_curious_01_head_angle_-20
anim_cubeconnection_connectionfailure_01
anim_blackjack_swipe_01
anim_gotosleep_getout_02
anim_explorer_idle_02_head_angle_30
anim_nowifi_trouble_icon_01
anim_qa_motors_backwards_500ms_01
anim_eyecontact_lookloop_01_head_angle_20
anim_message_getin_01
anim_rtpkeepaway_reacttocube_02
anim_blackjack_spread_01
anim_dizzy_reaction_hard_01
anim_feedback_iloveyou_02
anim_dancebeat_listening_01
anim_gotosleep_sleeping_02
anim_onboarding_cube_psychic_01
anim_rtshake_drive_loop_01
anim_slowbump_getinforward_01
anim_pounce_01
anim_keepaway_idle_01
anim_eyecontact_giggle_01_head_angle_40
anim_cubespinner_anticgamesuccess_01
anim_rtsound_oncharger_asleep_30right_01
anim_explorer_huh_01_head_angle_30
anim_keepaway_losehand_01
anim_explorer_center_left_01
anim_pounce_04
anim_heldonpalm_nestling_01
anim_cubespinner_sessionsuccess_01
anim_message_tts_getout
anim_rtshake_lv2rtonground_01
anim_rtsound_oncharger_observe_150left_01
anim_gif_no_01
anim_photo_shutter_01
anim_keepaway_pounce_03
anim_reacttocliff_edge_01
anim_rtsound_oncharger_observe_right_01
anim_keepaway_getready_01
anim_rtshake_lv1loop_01
anim_message_rewind_01
anim_gazing_lookatfaces_getin_left_01
anim_lookinplaceforfaces_keepalive_long_04_head_angle_20
anim_rtshake_drive_getout_01
anim_onboarding_eyecontact_loop_01
anim_greeting_goodmorning_01
anim_weather_stars_01
anim_slowpoke_getin_03
anim_freeplay_reacttoface_sayname_02_head_angle_20
anim_slowbump_getinbackward_01
anim_movement_left_01
anim_reacttocliff_stuckonedge_getin_01
anim_eyecontact_getout_01_head_angle_20
anim_petting_lvl2_getout_01
anim_reacttoblock_react_01
anim_pairing_icon_update
anim_referencing_giggle_01_head_angle_40
anim_qa_sync_3_30_30
anim_greeting_happy_03_head_angle_40
anim_movement_backward_01
anim_reacttocliff_stop_02
anim_getout_tts_02
anim_hiking_driving_start_01
anim_movement_lookinplaceforfaces_long_head_angle_-20
anim_rtshake_getin_01
anim_laser_drive_03
anim_gazing_lookatfaces_getin_left_01_head_angle_20
anim_reacttocliff_reaction_rear_right_01
anim_launch_startsleeping_01
anim_keepaway_fakeout_05
anim_reacttoblock_happydetermined_02
anim_pounce_drive_02
anim_cube_drive_getout_01
anim_eyepose_sad_down01
anim_reacttocliff_turtleroll_06
anim_keepaway_idleliftdown_02
anim_reacttocliff_turnright_180_01
anim_cubedocking_drive_loop_01
anim_referencing_squint_02_head_angle_-20
anim_eyecontact_right_contact_01_head_angle_40
anim_keepaway_backup_01
anim_wakeword_turn_left_135degrees_01
anim_rtmotion_right_getout_01
anim_wakeword_turn_right_45degrees_01
anim_reacttocliff_turtlerollfail_01
anim_eyecontact_lookloop_01
anim_rtpkeepaway_idle_03
anim_rtshake_lv1onground_loosepixel_01
anim_timercancel_getin_01
anim_movement_reacttoface_01_head_angle_40
anim_pounce_drive_03
anim_observe_onchager_getin_01
anim_gotosleep_off_01
anim_eyecontact_giggle_01
anim_explorer_idle_02_head_angle_20
anim_reacttohabitat_subtle_01
anim_driving_upset_end_01
anim_eyecontact_smile_01_head_angle_-20
anim_qa_motors_head_down_500ms_01
anim_reacttocliff_stuckonedge_right_getin_01
anim_reacttocliff_stuckonedge_01
anim_blackjack_deal_01_temp_remap
anim_rtsound_oncharger_observe_120left_01
anim_reacttocliff_huh_01
anim_rtsound_oncharger_asleep_behind_right_01
anim_driving_upset_start_01
anim_dancebeat_idle_01
anim_qa_motors_90left_turn_500ms_01
anim_heldonpalm_edge_relaxed_01
anim_rtsound_offcharger_observe_30left_01
anim_message_announce_tonoone_01
anim_petting_bliss_getout_01
anim_keepaway_idleliftdown_01
anim_wakeword_loop_02
anim_hiking_driving_loop_02
anim_launch_wakeup_enddriving_01
anim_hiking_lookaround_01
anim_cubespinner_roundsuccess_01
anim_dancebeat_getin_01
anim_onboarding_cube_drive_getin_01
anim_movement_lookinplaceforfaces_medium
anim_gazing_lookatsurfaces_getin_right_01_head_angle_40
anim_slowpoke_getinforward_01
anim_movement_comehere_reaction_01
anim_cubeconnection_loop_04
anim_weather_cold_01
anim_dizzy_shake_stop_01
anim_dancebeat_getout_02
anim_weather_cloud_01_temp_remap
anim_keepaway_losegame_03
anim_eyecontact_right_thought_01
anim_lookinplaceforfaces_keepalive_long
anim_timerset_getin_02
anim_keepaway_miss_reaction_01
anim_wakeword_getout_groggyeyes_turnright_90degrees_01
anim_rtpickup_loop_10
anim_petting_lvl1_01
anim_dancebeat_headlift_01
anim_pounce_reacttoobj_01_shorter
anim_pounce_lookgetout_01
anim_dancebeat_headliftbody_right_large_01
anim_eyecontact_center_thought_01_head_angle_-20
anim_keepalive_eyesonly_loop_03
anim_lookatdevice_icon_end
anim_feedback_badrobot_01
anim_message_getout_01
anim_chargerdocking_requestwait_01
anim_eyepose_blink_high
anim_rtpickup_putdown_01
anim_onboarding_cube_drive_getout_01
anim_eyepose_sad_down
anim_feedback_meanwords_02
anim_launch_enddriving_01
anim_inspectheldcube_loop_02
anim_eyepose_scrutinizing
anim_rtsound_offcharger_asleep_30right_01
anim_rtshake_lv2rtonground_iconmess_01
anim_explorer_planning_getin_01
anim_lookinplaceforfaces_keepalive_medium_head_angle_20
anim_freeplay_reacttoface_sayname_01
anim_laser_drive_02
anim_eyecontact_lookloop_02_head_angle_20
anim_reacttoface_unidentified_01_head_angle_40
anim_rtsound_oncharger_observe_60right_01
anim_keepaway_getout_satisfied_01
anim_rtp_blackjack_playerno_01
anim_feedback_meanwords_01
anim_heldonpalm_putdown_nervous_01
anim_keepaway_idle_02
anim_blackjack_speech_01
anim_reacttocliff_reaction_front_left_01
anim_dancebeat_eyehold_01
anim_qa_intervals_30frame
anim_cube_drive_getin_01
anim_driving_upset_loop_02
anim_eyepose_bothered
anim_keepaway_losehand_02
anim_qa_lift_updown_alt
anim_movement_lookinplaceforfaces_medium_head_angle_40
anim_petdetection_reaction_dog_01
anim_gotosleep_sleeping_01
anim_explorer_idle_01_head_angle_40
anim_nowifi_icon_01
anim_keepaway_wingame_03
anim_eyecontact_squint_01_head_angle_-20
anim_wakeword_turn_behind_left_01
anim_cubespinner_gamefail_02
anim_keepaway_getout_engaged_01
anim_slowpoke_getin_01
anim_exploring_test_01
anim_keepaway_winhand_02
anim_meetvictor_sayname_01
anim_freeplay_reacttoface_sayname_03
anim_reacttocliff_stop_01
anim_exploring_intervals_4s
anim_meetvictor_getin_01_head_angle_20
anim_howold_slowloop_01
anim_reacttocliff_reaction_back_01
anim_onboarding_cube_drive_loop_01
anim_rtsound_oncharger_asleep_30left_01
anim_eyecontact_center_contact_01
anim_blackjack_victorpush_01
anim_movement_lookinplaceforfaces_medium_head_angle_20
anim_weather_rain_01_temp_remap
anim_chargerdocking_reaction_01
anim_weather_sunny_01_temp_remap
anim_cubespinner_anticgamefail_02
anim_gazing_lookatsurfaces_getin_right_01_head_angle_-20
anim_launch_wakeup_drivingloop_01
anim_movement_forward_01
anim_rtsound_oncharger_asleep_left_01
anim_keepaway_getout_03
anim_chargerdocking_comeoff_straight_03
anim_keepaway_pounce_05
anim_gotosleep_sleeping_04
anim_eyepose_worried
anim_lookatphone_loop_01
anim_pounce_lookloop_01
anim_dancebeat_listening_03
anim_pounce_fail_04
anim_feedback_goodrobot_01
anim_lookinplaceforfaces_keepalive_long_04_head_angle_-20
anim_chargerdocking_leftturn_01
anim_codelab_getin_01
anim_cubespinner_getin_01
anim_lookinplaceforfaces_keepalive_short_head_angle_-20
anim_explorer_huh_01_head_angle_10
anim_rtpickup_loop_04
anim_generic_look_up_idle_01
anim_rtsound_oncharger_asleep_120right_01
anim_reacttocliff_turtlerollfail_02
anim_meetvictor_lookface_loop_01
anim_weather_windy_01_temp_remap
anim_petting_getin_01
anim_chargerdocking_searchforchargerturns_02
anim_slowpoke_getin_04
anim_explorer_lookaround_01
anim_eyecolorreact_random
anim_cubespinner_loop_01
anim_onboarding_reacttoface_happy_01_head_angle_40
anim_eyepose_furious
anim_gazing_lookatsurfaces_getin_left_01_head_angle_-20
anim_rtshake_lv1onhand_loosepixel_01
anim_onboarding_lookdown_lookaround_01
anim_lookinplaceforfaces_keepalive_long_02_head_angle_20
anim_keepaway_pounce_02
anim_slowpoke_loop_02
anim_meetvictor_youarebob_01
anim_eyepose_default
anim_wakeword_turn_behind_right_01
anim_weather_snow_01_temp_remap
anim_rtsound_offcharger_asleep_front_01
anim_exploring_test_rup_rdown
anim_rtsound_oncharger_observe_120right_01
playpenFailAnim
anim_eyecontact_squint_01_head_angle_40
anim_rtsound_offcharger_asleep_120right_01
anim_gazing_lookatfaces_getin_right_01_head_angle_20
anim_cubeconnection_loop_02
anim_keepaway_pounce_01
anim_dizzy_reaction_medium_03
anim_keepaway_fakeout_02
anim_gotosleep_sleeping_03
anim_referencing_squint_02
anim_gif_gleeserious_01
anim_rtsound_oncharger_asleep_ambient_01
anim_launch_reacttocubepickup
anim_slowpoke_getout_03
anim_rtmotion_turn45left_01
anim_pounce_drive_01
anim_petdetection_reaction_dog_04
anim_qa_motors_90rightforward_arcturn_500ms_01
anim_rtp_blackjack_request_01
anim_power_offon_03
anim_chargerdocking_searchforchargerturns_01
anim_chargerdocking_rightturn_01
anim_rtsound_offcharger_observe_front_01
anim_exploring_test_fast_hardstop
anim_keepaway_pounce_04
anim_eyepose_hurt
anim_message_record_reaction_01
anim_dancebeat_scoot_left_01
anim_eyecontact_squint_01
anim_knowledgegraph_success_01
anim_dizzy_reaction_soft_02
anim_greeting_happy_03_head_angle_-20
anim_dancebeat_quit_01
anim_heldonpalm_relaxed_idle_01
anim_keepaway_idle_03
anim_timercancel_01
anim_pounce_03
anim_gazing_lookatvector_reaction_01_head_angle_40
anim_qa_reset_to_neutral
anim_pounce_long_01
anim_rtillumination_lighttodark_01
anim_movement_lookinplaceforfaces_long_head_angle_40
anim_eyecontact_left_contact_01_head_angle_40
anim_onboarding_wakeword_getout_01
anim_wakeword_turn_right_01
anim_dizzy_pickup_03
anim_volume_stage_min
anim_knowledgegraph_answer_01
anim_volume_stage_05
anim_chargerdocking_comeoff_straight_01
anim_rtsound_offcharger_asleep_behind_left_01
anim_slowpoke_loop_03
anim_rtpickup_putdown_04
anim_greeting_happy_03_head_angle_20
anim_eyecontact_left_contact_01_head_angle_-20
anim_gazing_lookatsurfaces_turn_left_01
anim_wakeword_idk_03
anim_referencing_smile_01_head_angle_40
anim_wakeword_turn_left_45degrees_01
anim_eyecontact_giggle_01_head_angle_-20
anim_turn_left_01
anim_wakeword_idk_01
anim_petdetection_reaction_cat_04
anim_keepaway_getout_02
anim_uperformance_binaryeyes_01
anim_inspectheldcube_loop_05
anim_lookatphone_getout_01
anim_keepaway_getin_02
anim_reacttocliff_stuckonedge_alert_01
anim_eyepose_concerned
anim_launch_startdriving_01
anim_rtshake_lv3pregetout_01
anim_lookatphone_getin_01
anim_reacttoface_unidentified_03_head_angle_-20
anim_qa_head_updown
anim_keepaway_getout_01
anim_reacttocliff_pickup_06
anim_slowpoke_getout_02
anim_rtsound_offcharger_asleep_30left_01
anim_dizzy_pickup_02
anim_keepaway_driving_loop_03
anim_eyecolorreact_switch_02
anim_slowpoke_getout_01
anim_lookinplaceforfaces_keepalive_long_head_angle_20
anim_reacttocliff_turnright_60_01
*/

// Displays a list of the available animations, or nil in case of failure

func LoadAnimationList() []string {
	var items []string
	animationList, err := Robot.Conn.ListAnimations(
		ctx,
		&vectorpb.ListAnimationsRequest{},
	)
	if err == nil {
		items = make([]string, len(animationList.AnimationNames))
		for i := 0; i < len(animationList.AnimationNames); i++ {
			items[i] = animationList.AnimationNames[i].Name
		}
	}
	return items
}

// Displays animation

func PlayAnimation(anim string, loops uint32, ignoreBodyTrack bool, ignoreHeadTrack bool, ignoreLiftTrack bool) {
	var a = &vectorpb.Animation{
		Name: anim, // ignore SSL warnings
	}
	_, _ = Robot.Conn.PlayAnimation(
		ctx,
		&vectorpb.PlayAnimationRequest{
			Animation:       a,
			Loops:           loops,
			IgnoreBodyTrack: ignoreBodyTrack,
			IgnoreHeadTrack: ignoreHeadTrack,
			IgnoreLiftTrack: ignoreLiftTrack,
		},
	)
}
