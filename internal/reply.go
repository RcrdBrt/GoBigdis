/*
	GoBigdis is a persistent database that implements the Redis server protocol.
    Copyright (C) 2021  Riccardo Berto

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.
*/
package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type ReplyWriter io.WriterTo

type StatusReply struct {
	Code string
}

func (r *StatusReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte("+" + r.Code + "\r\n"))
	return int64(n), err
}

type IntegerReply struct {
	number int
}

func (r *IntegerReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(":" + strconv.Itoa(r.number) + "\r\n"))
	return int64(n), err
}

type ErrorReply struct {
	msg string
}

func (r *ErrorReply) WriteTo(w io.Writer) (int64, error) {
	n, err := fmt.Fprintf(w, "-%s", r.msg)

	return int64(n), err
}

type BulkReply struct {
	value []byte
}

func writeBytes(value interface{}, w io.Writer) (int64, error) {
	//it's a NullBulkReply
	if value == nil {
		n, err := w.Write([]byte("$-1\r\n"))
		return int64(n), err
	}
	switch v := value.(type) {
	case []interface{}:
		if len(v) == 0 {
			n, err := w.Write([]byte("*0\r\n"))
			return int64(n), err
		}
		wrote, err := writeMultiBytes(v, w)
		return int64(wrote), err

	case string:
		if len(v) == 0 {
			n, err := w.Write([]byte("$-1\r\n"))
			return int64(n), err
		}
		wrote, err := w.Write([]byte("$" + strconv.Itoa(len(v)) + "\r\n"))
		if err != nil {
			return int64(wrote), err
		}
		wroteBytes, err := w.Write([]byte(v))
		if err != nil {
			return int64(wrote + wroteBytes), err
		}
		wroteCrLf, err := w.Write([]byte("\r\n"))
		return int64(wrote + wroteBytes + wroteCrLf), err
	case []byte:
		if len(v) == 0 {
			n, err := w.Write([]byte("$-1\r\n"))
			return int64(n), err
		}
		wrote, err := w.Write([]byte("$" + strconv.Itoa(len(v)) + "\r\n"))
		if err != nil {
			return int64(wrote), err
		}
		wroteBytes, err := w.Write(v)
		if err != nil {
			return int64(wrote + wroteBytes), err
		}
		wroteCrLf, err := w.Write([]byte("\r\n"))
		return int64(wrote + wroteBytes + wroteCrLf), err
	case int:
		wrote, err := w.Write([]byte(":" + strconv.Itoa(v) + "\r\n"))
		if err != nil {
			return int64(wrote), err
		}
		return int64(wrote), err
	}

	Debugf("Invalid type sent to writeBytes: %v", reflect.TypeOf(value).Name())
	return 0, errors.New("invalid type sent to writeBytes")
}

func (r *BulkReply) WriteTo(w io.Writer) (int64, error) {
	return writeBytes(r.value, w)
}

type MonitorReply struct {
	c <-chan string
}

func (r *MonitorReply) WriteTo(w io.Writer) (int64, error) {
	statusReply := &StatusReply{}
	totalBytes := int64(0)
	for line := range r.c {
		statusReply.Code = line
		if n, err := statusReply.WriteTo(w); err != nil {
			totalBytes += n
			return int64(totalBytes), err
		} else {
			totalBytes += n
		}
	}
	return totalBytes, nil
}

//for nil reply in multi bulk just set []byte as nil
type MultiBulkReply struct {
	values []interface{}
}

func MultiBulkFromMap(m map[string]interface{}) *MultiBulkReply {
	values := make([]interface{}, len(m)*2)
	i := 0
	for key, val := range m {
		values[i] = []byte(key)
		switch v := val.(type) {
		case string:
			values[i+1] = []byte(v)
		default:
			values[i+1] = val

		}

		i += 2
	}
	return &MultiBulkReply{values: values}
}

func writeMultiBytes(values []interface{}, w io.Writer) (int64, error) {
	if values == nil {
		return 0, errors.New("nil in multi bulk replies are not ok")
	}
	wrote, err := w.Write([]byte("*" + strconv.Itoa(len(values)) + "\r\n"))
	if err != nil {
		return int64(wrote), err
	}
	wrote64 := int64(wrote)
	for _, v := range values {
		wroteBytes, err := writeBytes(v, w)
		if err != nil {
			return wrote64 + wroteBytes, err
		}
		wrote64 += wroteBytes
	}
	return wrote64, err
}

func (r *MultiBulkReply) WriteTo(w io.Writer) (int64, error) {
	return writeMultiBytes(r.values, w)
}

func ReplyToString(r ReplyWriter) (string, error) {
	var b bytes.Buffer

	_, err := r.WriteTo(&b)
	if err != nil {
		return "ERROR!", err
	}
	return b.String(), nil
}

type MultiChannelWriter struct {
	Chans []*ChannelWriter
}

func (c *MultiChannelWriter) WriteTo(w io.Writer) (n int64, err error) {
	chans := make(chan struct{}, len(c.Chans))
	for _, elem := range c.Chans {
		go func(elem io.WriterTo) {
			defer func() { chans <- struct{}{} }()
			if n2, err2 := elem.WriteTo(w); err2 != nil {
				n += n2
				err = err2
				return
			} else {
				n += n2
			}
		}(elem)
	}
	for i := 0; i < len(c.Chans); i++ {
		<-chans
	}
	return n, err
}

type ChannelWriter struct {
	FirstReply []interface{}
	Channel    chan []interface{}
	clientChan chan struct{}
}

func (c *ChannelWriter) WriteTo(w io.Writer) (int64, error) {
	totalBytes, err := writeMultiBytes(c.FirstReply, w)
	if err != nil {
		return totalBytes, err
	}

	for {
		select {
		case <-c.clientChan:
			return totalBytes, err
		case reply := <-c.Channel:
			if reply == nil {
				return totalBytes, nil
			} else {
				wroteBytes, err := writeMultiBytes(reply, w)
				// FIXME: obvious overflow here,
				// Just ignore? Who cares?
				totalBytes += wroteBytes
				if err != nil {
					return totalBytes, err
				}
			}
		}
	}
}
