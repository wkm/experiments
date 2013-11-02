----------------------------- MODULE hourclock ------------------------------
EXTENDS naturals
VARIABLE hr
HCini  == hr \in (1 .. 12)
HCnxt  == hr' = IF hr # 12 THEN hr + 1 ELSE 1
HC == HCini /\ [][HCnxt]_hr
-----------------------------------------------------------------------------
THEOREM HC => []HCini 
=============================================================================
\* Modification History
\* Last modified Wed Jun 19 13:46:22 EDT 2013 by wiktor
\* Created Wed Jun 19 13:44:06 EDT 2013 by wiktor
